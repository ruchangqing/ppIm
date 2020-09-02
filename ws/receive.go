package ws

import (
	"github.com/gin-gonic/gin"
	"ppIm/global"
	"ppIm/model"
)

// 已认证用户收到消息
func Receive(c *Connection, message Message) {
	switch message.Cmd {
	case 0:
		c.Conn.WriteJSON(WsMsg(0, 200, "test okay!", nil))
		break
	case 2:
		// 发私聊信息
		data := message.Data
		toUid := data["toUid"].(int)
		content := data["content"]
		var friendList model.FriendList
		var count int
		global.Mysql.Where("uid = ? and f_uid = ?", c.Uid, toUid).Select("id,uid,f_uid").First(&friendList).Count(&count)
		if count == 0 {
			c.Conn.WriteJSON(WsMsg(2, 0, "对方不是你的好友！", nil))
		} else {
			SendToUser(toUid, 2, 1, "有新消息", gin.H{
				"sendUid":     c.Uid,
				"content": content,
			})
		}
		break
	default:
		c.Conn.WriteJSON(WsMsg(-1, 0, "UnKnow message struct!", nil))
		break
	}
}
