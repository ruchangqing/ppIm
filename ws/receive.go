package ws

// 已认证用户收到消息
func Receive(c *Connection, message Message) {
	switch message.Cmd {
	case 0:
		c.Conn.WriteJSON(WsMsg(0, 200, "test okay!", nil))
		break
	default:
		c.Conn.WriteJSON(WsMsg(-1, 0, "UnKnow message struct!", nil))
		break
	}
}
