package ws

// 发送消息给uid
func SendToUid(uid uint, router string, code int, msg string, data interface{}) {
	client := Connections[uid]
	client.Conn.WriteJSON(WsMsg(router, code, msg, data))
}