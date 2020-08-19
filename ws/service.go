package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"ppIm/api"
	"sync"
)

// 连接结构体
type Connection struct {
	ClientId int
	Uid      uint
	Conn     *websocket.Conn
}

// 所有已认证连接结构体，key为用户id
var Connections = make(map[uint]Connection)

// 接收消息结构规范
type Message struct {
	Route string
	Data  map[string]interface{}
}

var Onlines = 0                 // 连接数
var OnlinesMutex sync.Mutex     // 连接数锁
var MaxClientId = 8880000000    // 目前最大生成的client_id
var GenClientIdMutex sync.Mutex // 生成client_id锁

// 给已认证用户生成client_id
func GenClientId() int {
	GenClientIdMutex.Lock()
	MaxClientId++
	temp := MaxClientId
	GenClientIdMutex.Unlock()
	return temp
}

// websocket状态接口
func Status(ctx *gin.Context) {
	api.R(ctx, 200, "status", gin.H{
		"connections": Connections,
		"clients":     MaxClientId - 8880000000,
		"onlines":     Onlines,
	})
}

// 发送消息给客户端规范
func WsMsg(router string, code int, msg string, data interface{}) gin.H {
	return gin.H{
		"route": router,
		"code":  code,
		"msg":   msg,
		"data":  data,
	}
}
