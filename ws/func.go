package ws

import (
	"github.com/gin-gonic/gin"
	"ppIm/global"
	"ppIm/servers"
)

// 发送消息给客户端规范
func WsMsg(cmd int, success int, msg string, data interface{}) gin.H {
	return gin.H{
		"cmd":  cmd,
		"code": success,
		"msg":  msg,
		"data": data,
	}
}

// 发送消息给uid
func SendToUser(uid int, cmd int, success int, msg string, data interface{}) {
	//if IsOnline(uid) {
	//	client := Connections[uid]
	//	client.Conn.WriteJSON(WsMsg(cmd, success, msg, data))
	//}
}

// 判断用户是否在线(本机)
func IsOnline(uid int) bool {
	if _, ok := UidToClientId[uid]; ok {
		return true
	} else {
		return false
	}
}

// 判断用户是否在线(集群)
func IsOnlineCluster(uid int) bool {
	isOnline := false
	for _, serverAddress := range servers.Servers {
		if serverAddress == global.ServerAddress {
			//调用本机方法查询uid在线
			if IsOnline(uid) {
				isOnline = true
				break
			}
		} else {
			//通过RPC调用其他集群查询uid在线
			if RpcIsOnline(serverAddress, uid) {
				isOnline = true
				break
			}
		}
	}
	return isOnline
}
