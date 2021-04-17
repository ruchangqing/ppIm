package ws

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
)

// 接收消息结构规范
type Message struct {
	Cmd  int
	Data map[string]interface{}
}

// websocket状态接口
func Status(ctx *gin.Context) {
	api.R(ctx, 200, "status", gin.H{
		"connections":   Connections,
		"uidToClientId": UidToClientId,
		"online":        len(Connections),
	})
}

// 发送消息给客户端规范
func WsMsg(cmd int, success int, msg string, data interface{}) gin.H {
	return gin.H{
		"cmd":  cmd,
		"code": success,
		"msg":  msg,
		"data": data,
	}
}
