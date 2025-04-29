package main

import "encoding/json"

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "New socket connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			manager.clients[conn] = false
			jsonMessage, _ := json.Marshal(&Message{Content: "A socket disconnected."})
			manager.send(jsonMessage, conn)
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}
