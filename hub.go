package main

// the type "hub" is a struct containing all the connections,
// a channel to broadcast some message
// a channel to register a new connection
// a channel to unregister a new connection
// a channel to pause the video
type hub struct {
	connections map[*connection]bool

	broadcast chan []byte

	register chan *connection

	unregister chan *connection

	pause chan bool
}

// newHub() returns a pointer to a hub
// creating all the required pieces of the struct
func newHub() *hub {
	return &hub {
		broadcast: make(chan []byte),
		register: make(chan *connection),
		unregister: make(chan *connection),
		connections: make(map[*connection]bool),
		pause: make(chan bool),
	}
}

// the method run on *hub loops forever (it is expected to do this on another goroutine
// then, it blocks
// on register, we push it to the pool of connections,
// on unregister, we remove the connection from the pool and close it's channel
// on broadcast, we simply run through all the connections and send it a message
// on pause, we send a special message that all videos should start playing
func (h *hub) run() {
	for {
		select {
			case c := <-h.register:
				h.connections[c] = true
			case c := <-h.unregister:
				if _, ok := h.connections[c]; ok {
					delete(h.connections, c)
					close(c.send)
				}
			case m := <-h.broadcast:
				for c := range h.connections {
					select {
						case c.send <- m:
						default:
							delete(h.connections, c)
							close(c.send)
					}
				}
			case p := <-h.pause:
				if p {
					h.broadcast <- []byte("@pause")
				} else {
					h.broadcast <- []byte("@unpause")
				}
		}
	}
}
