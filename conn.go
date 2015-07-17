package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// the connection type is a struct,
// contaning the websocket.Conn pointer, from the Gorilla Toolkit,
// a channel to send in a message,
// and the hub it belongs.
type connection struct {
	ws *websocket.Conn

	send chan []byte

	h *hub
}

// the reader() method loops reading messages from the 'ws', and sending it in the 'send' channel
// if the message is a @pause, sends 'true' to the pause channel of the hub
// if the message is a @unpause, sends false to the pause channel of the hub
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		messageStr := string(message)
		if messageStr == "@pause" {
			c.h.pause <- true
		} else if messageStr == "@unpause" {
			c.h.pause <- false
		} else {
			c.h.broadcast <- message
		}
	}
	c.ws.Close()
}

// the writer drains the send channel and writes it to the WebSocket
func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

// upgrader makes http.ResponseWriter and http.Request websocket aware
var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

// the wsHandler struct maintains a pointer to the hub, in order to handle the websocket connection
type wsHandler struct {
	h *hub
}

// ServeHTTP on wsHandler implements the Server interface from the stdlib
// upgrades its ResponseWriter and Request structs,
// creates a connection, registers the current socket,
// defers unregister to the end of this func,
// sends the writer to another goroutine to handle incoming data in it's channel,
// and blocks, waiting for messages
func (wsh wsHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	c := &connection{send: make(chan []byte, 256), ws: ws, h: wsh.h}
	c.h.register <- c
	defer func() {c.h.unregister <- c}()
	go c.writer()
	c.reader()
}
