package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"ppIm/middleware"
	"time"
)

// 升级http为websocket服务
var WebsocketUpgrade = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocket服务入口
func WebsocketEntry(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	var conn *websocket.Conn
	var err error
	conn, err = WebsocketUpgrade.Upgrade(w, r, nil) // 升级为websocket协议
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	var c Connection
	var counter int

	ClientCounterLocker.Lock()
	c.Conn = conn
	ClientCounter++
	counter = ClientCounter
	Connections[counter] = &c
	ClientCounterLocker.Unlock()

	c.ClientId = GenClientId(counter)

	// 15秒内没收到token绑定成功的断开连接
	time.AfterFunc(15*time.Second, func() {
		if c.Uid == 0 {
			conn.Close()
		}
	})

	// 必须死循环，gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// json消息转换
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			// 接收到非json数据
			fmt.Println("[Websocket]消息解析失败: " + string(msg))
			conn.WriteJSON(Message{-1, "消息格式错误"})
			continue
		}

		// bind绑定uid和client_id，这是必须绑定的才能通信的
		if message.MsgType == Sign {
			if c.Uid == 0 {
				jwtToken := message.MsgContent
				id, err := middleware.ParseToken(ctx, jwtToken)
				if err != "" {
					fmt.Println(err)
					conn.WriteJSON(Message{SignFail, "认证失败"})
					continue
				}
				c.Uid = id
				UidToClientId[c.Uid] = c.ClientId // 认证成功后注册到已认证连接表，方便查询对应clientId
				conn.WriteJSON(Message{SignSuccess, "认证成功"})
			}
		} else {
			if c.Uid > 0 {
				conn.WriteJSON(Message{OK, "发送成功"})
			} else {
				conn.WriteJSON(Message{Fail, "您还未认证"})
			}
		}
	}

	ClientCounterLocker.Lock()
	delete(Connections, counter)
	if c.Uid > 0 {
		delete(UidToClientId, c.Uid)
	}
	ClientCounterLocker.Unlock()
}
