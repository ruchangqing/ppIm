package ws

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	"ppIm/global"
	"strconv"
)

func StatusApi(ctx *gin.Context) {
	api.R(ctx, global.SUCCESS, "status", gin.H{
		"connections":   Connections,
		"uidToClientId": UidToClientId,
		"online":        len(Connections),
	})
}

func IsOnlineApi(ctx *gin.Context) {
	uid, _ := strconv.Atoi(ctx.DefaultPostForm("uid", "0"))
	api.R(ctx, global.SUCCESS, "isOnline", gin.H{
		"isOnline":        LocalIsOnline(uid),
		"isOnlineCluster": IsOnline(uid),
	})
}

// 发送给所有客户端测试性能
func SendToAll(ctx *gin.Context) {
	for _, v := range Connections {
		v.Conn.WriteJSON(Message{Cmd: -1, Body: "测试消息收发"})
	}
}