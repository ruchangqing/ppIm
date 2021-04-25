package ws

import (
	"ppIm/global"
	"ppIm/servers"
	"strconv"
	"strings"
)

func LocalIsOnline(uid int) bool {
	if _, ok := UidToClientId[uid]; ok {
		return true
	} else {
		return false
	}
}

func LocalSendToUser(uid int, message Message) {
	clientId := UidToClientId[uid]
	arr := strings.Split(clientId, "@@")
	number, _ := strconv.Atoi(arr[1])
	Connections[number].Conn.WriteJSON(message)
}

// 判断用户是否在线
func IsOnline(uid int) bool {
	isOnline := false
	for _, serverAddress := range servers.Servers {
		if serverAddress == global.ServerAddress {
			//调用本机方法查询uid在线
			if LocalIsOnline(uid) {
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

// 发送消息给用户
func SendToUser(uid int, message Message) {
	go func() {
		for _, serverAddress := range servers.Servers {
			if serverAddress == global.ServerAddress {
				//调用本机方法查询uid在线
				if LocalIsOnline(uid) {
					LocalSendToUser(uid, message)
					break
				}
			} else {
				//通过RPC调用其他集群查询uid在线
				if RpcIsOnline(serverAddress, uid) {
					RpcSendToUser(serverAddress, message)
					break
				}
			}
		}
	}()
}

// 发送消息给群组
func SendToGroup(userIdList []int, message Message) {
	go func() {
		for _, uid := range userIdList {
			SendToUser(uid, message)
		}
	}()
}
