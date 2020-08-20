package ws

import "github.com/gorilla/websocket"

// 已认证用户收到消息
func Receive(conn *websocket.Conn, message Message) {
	switch message.Route {
	case "test":
		conn.WriteJSON(WsMsg("error", 200, "test okay!", nil))
		break
	default:
		conn.WriteJSON(WsMsg("error", 500, "UnKnow message struct!", nil))
		break
	}
}
