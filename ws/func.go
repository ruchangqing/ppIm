package ws

// 发送消息给uid
func SendToUser(uid int, cmd int, success int, msg string, data interface{}) {
	//if IsOnline(uid) {
	//	client := Connections[uid]
	//	client.Conn.WriteJSON(WsMsg(cmd, success, msg, data))
	//}
}

// 判断用户是否在线
func IsOnline(uid int)bool {
	//if _, ok := Connections[uid]; ok {
	//	return true
	//} else {
		return false
	//}
}