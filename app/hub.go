package app

import (
	"encoding/json"
)

var Hub = &clients{
	clients:    make(map[*Client]bool),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	broadcast:  make(chan []byte),
}

func (h *clients) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
			c.Message.Type = "init"
			c.Message.IP = c.Socket.RemoteAddr().String()
			c.Message.UserList = user_list
			msg_data, _ := json.Marshal(c.Message)
			c.Send <- msg_data
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.Send)
			}
		case data := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.Send <- data:
				default:
					delete(h.clients, c)
					close(c.Send)
				}
			}
		}
	}
}
