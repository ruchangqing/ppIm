package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var WsUpgrade = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	var conn *websocket.Conn
	var err error
	conn, err = WsUpgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	// 必须死循环，gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁
	for {
		t, reply, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// todo：业务操作
		fmt.Println(t, reply)
	}
}
