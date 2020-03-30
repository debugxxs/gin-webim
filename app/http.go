package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 将 http 请求升级为 websocket 请求
var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	user_list = []map[string]string{}
)

func WsServer(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("出错咯～", err)
	}
	// register：客户端连接，定义客户端基本结构
	client := &Client{Socket: ws, Send: make(chan []byte), Message: &Message{}}
	Hub.register <- client
	go client.Writer()
	client.Reader()

	// 当前 wsServer 结束时，关闭此客户端连接
	defer func() {
		client.Message.Type = "logout"
		user_list = remove(client.Message.ID)
		client.Message.UserList = user_list
		msg_data, _ := json.Marshal(client.Message)
		Hub.broadcast <- msg_data /* 将客户端断开的时间写入广播的数据管道 */
		Hub.unregister <- client  /* 将此客户端连接写入断开连接的数据管道 */
	}()
}
