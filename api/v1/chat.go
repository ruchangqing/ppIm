package v1

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	"ppIm/global"
	"ppIm/model"
	"ppIm/ws"
	"strconv"
	"time"
)

// 发送私信给好友
func SendMessageToUser(ctx *gin.Context) {
	toUid, _ := strconv.Atoi(ctx.PostForm("to_uid"))
	content := ctx.PostForm("content")
	messageType := 1 // 目前只支持消息类型文字1
	uid := int(ctx.MustGet("id").(float64))

	var friendList model.FriendList
	var count int
	global.Mysql.Where("uid = ? and f_uid = ?", uid, toUid).Select("id,uid,f_uid").First(&friendList).Count(&count)
	if count == 0 {
		api.R(ctx, 500, "对方不是你的好友", gin.H{})
	} else {
		var user model.User
		global.Mysql.Where("id = ?", uid).First(&user)
		now := time.Now().Format("2006-01-02 15:04:05")
		ws.SendToUser(toUid, 2, 1, "有新消息", gin.H{
			"sender": gin.H{
				"uid":      uid,
				"nickname": user.Nickname,
				"avatar":   user.Avatar,
			},
			"message": gin.H{
				"messageType": messageType,
				"content":     content,
				"created_at":  now,
			},
		})
		// 持久化消息记录
		chatUser := model.ChatUser{SendUid: uid, RecvUid: toUid, MessageType: messageType, Content: content, Status: 0, CreatedAt: now}
		global.Mysql.Create(&chatUser)
		api.Rt(ctx, 200, "发送成功", gin.H{})
	}
}
