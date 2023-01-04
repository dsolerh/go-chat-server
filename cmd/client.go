package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
)

// client represents a single chating user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan *message
	// room is the room this client is chating in.
	room *room
	// the user information
	userData objx.Map
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading:", err)
			return
		}
		msg.Timestamp = time.Now()
		msg.Name = c.userData["name"].(string)
		if avatarURL, ok := c.userData["avatar_url"]; ok {
			msg.AvatarURL = avatarURL.(string)
		}
		c.room.forward <- &msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			log.Println("Error writing:", err)
			return
		}
	}
}
