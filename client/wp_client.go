package client

import (
	"github.com/Rhymen/go-whatsapp"
)

type ChatClient interface {
	CreateConnection() error
	SendMessage(message, number string) error
}

type WpClient struct {
	connection whatsapp.Conn
}
