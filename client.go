package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			break
		}

		jsonMessage, _ := json.Marshal(&Message{Content: string(message), Sender: c.id})
		rdb.Publish(ctx, channel, jsonMessage)
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	sub := rdb.Subscribe(ctx, channel)
	defer sub.Close()
	ch := sub.Channel()

	for {
		select {
		case msg := <-ch:
			if err := c.socket.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				fmt.Println("Write error:", err)
				return
			}
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
