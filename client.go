package main

import (
	"github.com/gorilla/websocket"
)

//
type client struct {
	//socket is the websocket for Client
	socket *websocket.Conn

	//receive is a channel to receive message from other client
	receive chan []byte

	//room is the room for client chatting each other
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
