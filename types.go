package main

import "github.com/gorilla/websocket"

type ClientManager struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
}

type Client struct {
	id     string
	send   chan []byte
	socket *websocket.Conn
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}
