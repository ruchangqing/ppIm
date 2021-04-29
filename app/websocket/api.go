package websocket

import (
	"github.com/gin-gonic/gin"
	"ppIm/app/api"
	"ppIm/global"
)

// websocket本机服务状态
func StatusApi(ctx *gin.Context) {
	api.R(ctx, global.ApiSuccess, "status", gin.H{
		"connections":   Connections,
		"uidToClientId": UidToClientId,
		"online":        len(Connections),
	})
}

// 发送给所有客户端测试性能
func SendToAll(ctx *gin.Context) {
	for _, v := range Connections {
		v.Conn.WriteJSON(Message{Cmd: -1, Body: "测试消息收发"})
	}
}