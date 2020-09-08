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
	if content == "" {
		api.R(ctx, 500, "请输入内容", gin.H{})
		return
	}
	messageType := 1 // 目前只支持消息类型文字1
	uid := int(ctx.MustGet("id").(float64))

	var friendList model.FriendList
	var count int
	global.Mysql.Where("uid = ? and f_uid = ?", uid, toUid).Select("id,uid,f_uid").First(&friendList).Count(&count)
	if count == 0 {
		api.R(ctx, 500, "对方不是你的好友", gin.H{})
		return
	} else {
		var user model.User
		global.Mysql.Where("id = ?", uid).First(&user)
		now := time.Now().Unix()
		// 持久化消息记录
		chatUser := model.ChatUser{SendUid: uid, RecvUid: toUid, MessageType: messageType, Content: content, Status: 0, CreatedAt: now}
		global.Mysql.Create(&chatUser)
		ws.SendToUser(toUid, 2, 1, "有新消息", gin.H{
			"sender": gin.H{
				"uid":      uid,
				"nickname": user.Nickname,
				"avatar":   user.Avatar,
			},
			"message": gin.H{
				"id":          chatUser.Id,
				"messageType": messageType,
				"content":     content,
				"created_at":  now,
			},
		})
		api.Rt(ctx, 200, "发送成功", gin.H{})
	}
}

// 撤回私聊消息
func WithdrawMessageFromUser(ctx *gin.Context) {
	messageId, _ := strconv.Atoi(ctx.PostForm("message_id"))
	uid := int(ctx.MustGet("id").(float64))
	//查询消息记录
	var chatUser model.ChatUser
	global.Mysql.Where("id = ? and send_uid = ? and status <> ?", messageId, uid, -1).First(&chatUser)
	if chatUser.Id == 0 {
		api.R(ctx, 500, "消息不存在", gin.H{})
		return
	}
	now := time.Now().Unix()
	// 判断消息发送是否超过2分钟
	if now > chatUser.CreatedAt+120 {
		api.R(ctx, 500, "消息超过2分钟无法撤回", gin.H{})
		return
	}
	// 把消息状态改为已撤回
	global.Mysql.Model(&chatUser).Updates(map[string]interface{}{"status": -1})
	// 通过对方撤回消息
	ws.SendToUser(chatUser.RecvUid, 3, 1, "有新撤回消息", gin.H{
		"f_uid":     uid,
		"messageId": chatUser.Id,
	})
	api.Rt(ctx, 200, "撤回成功", gin.H{})
}
