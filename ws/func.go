package ws

// 发送消息给uid
func SendToUser(uid uint, router string, code int, msg string, data interface{}) {
	if isOnline(uid) {
		client := Connections[uid]
		client.Conn.WriteJSON(WsMsg(router, code, msg, data))
	}
}

// 判断用户是否在线
func isOnline(uid uint)bool {
	if _, ok := Connections[uid]; ok {
		return true
	} else {
		return false
	}
}