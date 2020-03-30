package app

import "github.com/gorilla/websocket"

type clients struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

// 单个socket客户端属性
type Client struct {
	Socket  *websocket.Conn
	Send    chan []byte
	Message *Message
}

// 数据传输的数据结构
type Message struct {
	ID       string              `json:"uuid"`
	IP       string              `json:"ip"`
	To       string              `json:"touuid"`
	Type     string              `json:"type"`
	User     string              `json:"username"`
	Content  string              `json:"content"`
	UserList []map[string]string `json:"user_list"`
}
