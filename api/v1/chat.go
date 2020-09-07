package v1

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	"ppIm/global"
	"ppIm/model"
	"ppIm/ws"
	"strconv"
)

// 发送私信给好友
func SendMessageToUser(ctx *gin.Context) {
	toUid, _ := strconv.Atoi(ctx.PostForm("username"))
	content := ctx.PostForm("message")
	messageType := 1 // 目前只支持消息类型文字1
	uid := int(ctx.MustGet("id").(float64))

	var friendList model.FriendList
	var count int
	global.Mysql.Where("uid = ? and f_uid = ?", uid, toUid).Select("id,uid,f_uid").First(&friendList).Count(&count)
	if count == 0 {
		api.R(ctx, -1, "对方不是你的好友", gin.H{})
	} else {
		var user model.User
		global.Mysql.Where("id = ?", uid).Select("id,nickname,avatar").First(&user)
		ws.SendToUser(toUid, 2, 1, "有新消息", gin.H{
			"sender": gin.H{
				"uid": uid,
				"nickname": user.Nickname,
				"avatar": user.Avatar,
			},
			"message": gin.H{
				"messageType": messageType,
				"content": content,
			},
		})
	}
}
