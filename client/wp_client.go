package client

import (
	"encoding/gob"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"os"
	"sync"
	"time"
)

type ChatClient interface {
	CreateConnection() error
	SendMessage(message, number string) error
}

type WpClient struct {
	connection *whatsapp.Conn
}

var (
	mx    sync.Mutex
	conns = make(map[int]*whatsapp.Conn)
)

func (w WpClient) CreateConnection() error {
	// Start Connection
	wac, err := whatsapp.NewConn(5 * time.Second)

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
			fmt.Fprintf(os.Stderr, "restoring failed: %v\n", err)
			return err
		}
	} else {
		// No saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
		}
	}

	// Save session
	err = writeSession(session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error saving session: %v\n", err)
	}

	<-time.After(1 * time.Second)

	fmt.Printf("login successful, session: %v\n", session)
	addConnPath(wac)
	//contact := wac.Store.Contacts
	//fmt.Println(contact)
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
	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: number + "@s.whatsapp.net",
		},
		ContextInfo: whatsapp.ContextInfo{
			QuotedMessageID: "",
		},
		Text: message,
	}

	msgId, err := wac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v", err)
		return err
	} else {
		fmt.Println("Message Sent -> ID : " + msgId)
		return nil
	}
}

func writeSession(session whatsapp.Session) error {
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

func addConnPath(conn *whatsapp.Conn) int {
	mx.Lock()
	var connId int = len(conns)
	conns[connId] = conn
	mx.Unlock()
	return connId
}

func getConn(connId int) *whatsapp.Conn {
	mx.Lock()
	var conn *whatsapp.Conn = conns[connId]
	mx.Unlock()
	return conn
}
