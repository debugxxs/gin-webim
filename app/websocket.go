package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (c *Client) Reader() {
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Hub.unregister <- c
			break
		}
		_ = json.Unmarshal(message, &c.Message)
		handleMsgType(c)
	}
}

func (c *Client) Writer() {
	for message := range c.Send {
		if c.Message.Type == "private" {
			if c.Message.To == c.Message.ID {
				_ = c.Socket.WriteMessage(websocket.TextMessage, message)
			}
		} else {
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
	_ = c.Socket.Close()
}

func handleMsgType(c *Client) {
	switch c.Message.Type {
	case "login":
		user_list = inert(c.Message.ID, c.Message.User)
		c.Message.UserList = user_list
		msg_data, _ := json.Marshal(c.Message)
		Hub.broadcast <- msg_data
	case "message":
		msg_data, _ := json.Marshal(c.Message)
		Hub.broadcast <- msg_data
	case "private":
		msg_data, _ := json.Marshal(c.Message)
		Hub.broadcast <- msg_data
	case "logout":
		c.Message.Type = "logout"
		user_list = remove(c.Message.ID)
		c.Message.UserList = user_list
		msg_data, _ := json.Marshal(c.Message)
		Hub.broadcast <- msg_data
		Hub.unregister <- c
	default:
		fmt.Println("=======================")
	}
}

// 将新加入的成员写入user_list
func inert(uuid, user string) []map[string]string {
	user_list = append(user_list, map[string]string{
		"uuid":     uuid,
		"username": user,
	})
	return user_list
}

// 将断线的成员移除
func remove(uuid string) []map[string]string {
	var new_user_list = []map[string]string{}
	for _, item := range user_list {
		if item["uuid"] != uuid {
			fmt.Println(item["uuid"], uuid)
			new_user_list = append(new_user_list, item)
		}
	}
	return new_user_list
}
