package ws

import (
	"ppIm/global"
	"ppIm/servers"
)

// 发送消息给uid
func SendToUser(uid int, cmd int, success int, msg string, data interface{}) {
	//if IsOnline(uid) {
	//	client := Connections[uid]
	//	client.Conn.WriteJSON(WsMsg(cmd, success, msg, data))
	//}
}

// 判断用户是否在线
func IsOnline(uid int) bool {
	for _, serverAddress := range servers.Servers {
		if serverAddress == global.ServerAddress {
			//调用本机方法查询uid在线
		} else {
			//通过RPC调用其他集群查询uid在线
		}
	}
	return true
}
