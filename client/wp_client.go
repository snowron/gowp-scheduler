package client

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
)

type ChatClient interface {
	CreateConnection() error
	SendMessage(message, number string) error
}

type WpClient struct{}

var (
	mx    sync.Mutex
	conns = make(map[int]*whatsapp.Conn)
)

func (w WpClient) CreateConnection() error {
	// Start Connection
	const waitTime = 5
	wac, err := whatsapp.NewConnWithOptions(&whatsapp.Options{
		Timeout:         waitTime * time.Second,
		ShortClientName: "gowp-scheduler",
		LongClientName:  "gowp-scheduler client",
	})

	// nolint:gomnd
	wac.SetClientVersion(2, 2123, 7)
	if err != nil {
		panic(err)
	}

	// Load saved session
	session, err := readSession()

	if err == nil {
		// Restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return err
		}
	} else {
		// No saved session -> regular login
		err = regularLogin(wac)
		if err != nil {
			return err
		}
	}

	// Save session
	err = writeSession(&session)
	if err != nil {
		return err
	}

	<-time.After(1 * time.Second)

	fmt.Println("login successfully")
	addConn(wac)

	return err
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func (w WpClient) SendMessage(message, number string) error {
	wac := getConn(0)

	// Load saved session every send message called for unblock user web interface experience
	session, err := readSession()

	if err == nil {
		// Restore session
		_, _ = wac.RestoreWithSession(session)
	}

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: number + "@s.whatsapp.net",
		},
		ContextInfo: whatsapp.ContextInfo{
			QuotedMessageID: "",
		},
		Text: message,
	}

	msgID, err := wac.Send(msg)
	if err != nil {
		return err
	}

	fmt.Println("Message Sent -> ID : " + msgID)
	return nil
}

func writeSession(session *whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}
func regularLogin(wac *whatsapp.Conn) error {
	qr := make(chan string)

	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	_, err := wac.Login(qr)

	if err != nil {
		return err
	}
	return nil
}
func addConn(conn *whatsapp.Conn) {
	mx.Lock()
	var connID = len(conns)
	conns[connID] = conn
	mx.Unlock()
}

func getConn(connID int) *whatsapp.Conn {
	mx.Lock()
	var conn = conns[connID]
	mx.Unlock()
	return conn
}
