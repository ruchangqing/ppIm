package ws

import (
	"ppIm/global"
	"ppIm/servers"
	"strconv"
	"strings"
)

// 发送消息给uid
func SendToUser(uid int, msgType int, msgContent string) {
	message := Message{msgType, msgContent}
	for _, serverAddress := range servers.Servers {
		if serverAddress == global.ServerAddress {
			//调用本机方法查询uid在线
			if IsOnline(uid) {
				clientId := UidToClientId[uid]
				arr := strings.Split(string(clientId), "@@")
				number, _ := strconv.Atoi(arr[1])
				Connections[number].Conn.WriteJSON(message)
				break
			}
		} else {
			//通过RPC调用其他集群查询uid在线
			if RpcIsOnline(serverAddress, uid) {
				RpcSendToUser(serverAddress, uid, msgType, msgContent)
				break
			}
		}
	}
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
