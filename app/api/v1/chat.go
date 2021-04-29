package v1

import (
	"github.com/gin-gonic/gin"
	"ppIm/app/api"
	"ppIm/app/model"
	"ppIm/app/websocket"
	"ppIm/global"
	"strconv"
	"time"
)

type chat struct{}

var Chat chat

// 发送好友消息
func (chat) SendToUser(ctx *gin.Context) {
	toUid, _ := strconv.Atoi(ctx.PostForm("to_uid"))
	messageType := 1 // 目前只支持消息类型文字1
	body := ctx.PostForm("body")
	if body == "" {
		api.R(ctx, global.ApiFail, "请输入内容", gin.H{})
		return
	}
	uid := int(ctx.MustGet("id").(float64))

	var friendList model.FriendList
	var count int
	global.Db.Where("uid = ? and f_uid = ?", uid, toUid).Select("id,uid,f_uid").First(&friendList).Count(&count)
	if count == 0 {
		api.R(ctx, global.ApiFail, "对方不是你的好友", gin.H{})
		return
	} else {
		// 持久化消息记录
		chatMessage := model.ChatMessage{
			FromId:    uid,
			ToId:      toUid,
			Ope:       0,
			Type:      messageType,
			Body:      body,
			Status:    0,
			CreatedAt: time.Now().Unix(),
		}
		global.Db.Create(&chatMessage)
		// 发送实时消息
		message := websocket.Message{
			Cmd:    websocket.CmdReceiveFriendMessage,
			FromId: uid,
			ToId:   toUid,
			Ope:    0,
			Type:   messageType,
			Body:   body,
		}
		websocket.SendToUser(toUid, message)

		api.Rt(ctx, global.ApiSuccess, "发送成功", gin.H{"id": chatMessage.Id})
	}
}

// 撤回好友消息
func (chat) WithdrawFromUser(ctx *gin.Context) {
	messageId, _ := strconv.Atoi(ctx.PostForm("message_id"))
	uid := int(ctx.MustGet("id").(float64))
	//查询消息记录
	var chatMessage model.ChatMessage
	global.Db.Where("id = ? AND from_id = ? AND status <> ? AND ope = 0", messageId, uid, -1).First(&chatMessage)
	if chatMessage.Id == 0 {
		api.R(ctx, global.ApiFail, "消息不存在", gin.H{})
		return
	}
	now := time.Now().Unix()
	// 判断消息发送是否超过2分钟
	if now > chatMessage.CreatedAt+120 {
		api.R(ctx, global.ApiFail, "消息超过2分钟无法撤回", gin.H{})
		return
	}
	// 把消息状态改为已撤回
	global.Db.Model(&chatMessage).Updates(map[string]interface{}{"status": -1})
	// 通知对方撤回消息
	message := websocket.Message{
		Cmd:    websocket.CmdWithdrawFriendMessage,
		FromId: uid,
		ToId:   chatMessage.ToId,
		Ope:    0,
		Type:   0,
		Body:   strconv.Itoa(chatMessage.Id),
	}
	websocket.SendToUser(chatMessage.ToId, message)

	api.Rt(ctx, global.ApiSuccess, "撤回成功", gin.H{})
}

// 发送群消息
func (chat) SendToGroup(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	groupId, _ := strconv.Atoi(ctx.PostForm("group_id"))
	body := ctx.PostForm("body")
	if body == "" {
		api.R(ctx, global.ApiFail, "请输入内容", gin.H{})
		return
	}
	var group model.Group
	global.Db.Where("id = ?", groupId).First(&group)
	if group.Id == 0 {
		api.R(ctx, global.ApiFail, "群组不存在", gin.H{})
		return
	}
	var groupUser model.GroupUser
	global.Db.Where("group_id = ? AND user_id = ?", groupId, uid).First(&groupUser)
	if groupUser.Id == 0 {
		api.R(ctx, global.ApiFail, "您不是群成员", gin.H{})
		return
	}

	// 持久化消息记录
	messageType := 1
	chatGroup := model.ChatMessage{
		FromId:    uid,
		ToId:      groupId,
		Ope:       1,
		Type:      messageType,
		Body:      body,
		CreatedAt: time.Now().Unix(),
		Status:    0,
	}
	global.Db.Create(&chatGroup)

	// 发送实时消息
	type result struct {
		UserId int
	}
	var groupUserList []result
	global.Db.Raw("SELECT user_id FROM `group_user` WHERE group_id = >", groupId).Scan(&groupUserList)
	if len(groupUserList) > 0 {
		var userIdList []int
		for _, groupUser := range groupUserList {
			userIdList = append(userIdList, groupUser.UserId)
		}
		message := websocket.Message{
			Cmd:    websocket.CmdReceiveGroupMessage,
			FromId: uid,
			ToId:   groupId,
			Ope:    1,
			Type:   messageType,
			Body:   body,
		}
		websocket.SendToGroup(userIdList, message)
	}

	api.Rt(ctx, global.ApiSuccess, "发送成功", gin.H{"messageId": chatGroup.Id})
}

// 撤回群消息
func (chat) WithdrawFromGroup(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	messageId, _ := strconv.Atoi(ctx.PostForm("message_id"))
	//查询消息记录
	var chatMessage model.ChatMessage
	global.Db.Where("id = ? AND from_id = ? AND status <> ? AND ope = 1", messageId, uid, -1).First(&chatMessage)
	if chatMessage.Id == 0 {
		api.R(ctx, global.ApiFail, "消息不存在", gin.H{})
		return
	}
	now := time.Now().Unix()
	// 判断消息发送是否超过2分钟
	if now > chatMessage.CreatedAt+120 {
		api.R(ctx, global.ApiFail, "消息超过2分钟无法撤回", gin.H{})
		return
	}
	// 把消息状态改为已撤回
	global.Db.Model(&chatMessage).Updates(map[string]interface{}{"status": -1})
	// 发送实时消息
	type result struct {
		UserId int
	}
	var groupUserList []result
	global.Db.Raw("SELECT user_id FROM `group_user` WHERE group_id = >", chatMessage.ToId).Scan(&groupUserList)
	if len(groupUserList) > 0 {
		var userIdList []int
		for _, groupUser := range groupUserList {
			userIdList = append(userIdList, groupUser.UserId)
		}
		message := websocket.Message{
			Cmd:    websocket.CmdWithdrawGroupMessage,
			FromId: uid,
			ToId:   chatMessage.ToId,
			Ope:    1,
			Type:   0,
			Body:   strconv.Itoa(chatMessage.Id),
		}
		websocket.SendToGroup(userIdList, message)
	}

	api.Rt(ctx, global.ApiSuccess, "撤回成功", gin.H{})
}
