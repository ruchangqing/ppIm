package v1

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	"ppIm/global"
	"ppIm/model"
	"strconv"
	"time"
)

// 添加好友
func AddFriend(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	username := ctx.PostForm("username")
	channel := ctx.PostForm("channel")
	reason := ctx.PostForm("reason")

	var user model.User
	var count int
	global.Mysql.Where("username = ?", username).Select("id,username").First(&user).Count(&count)
	if count < 1 {
		api.R(ctx, 500, "用户未找到", nil)
		return
	}
	if user.Id == uid {
		api.R(ctx, 500, "不能添加自己", nil)
		return
	}

	var friendList model.FriendList
	global.Mysql.Where("uid = ? and f_uid = ?", uid, user.Id).Select("id,uid,f_uid").First(&friendList).Count(&count)
	if count > 0 {
		api.R(ctx, 500, "对方已经是好友", nil)
		return
	}

	var friendAdd model.FriendAdd
	global.Mysql.Where("uid = ? and f_uid = ? and status = 0", uid, user.Id).Select("id,uid,f_uid").First(&friendAdd).Count(&count)
	if count > 0 {
		api.R(ctx, 500, "请等待好友同意", nil)
		return
	}

	friendAdd.Uid = uid
	friendAdd.FUid = user.Id
	friendAdd.Reason = reason
	friendAdd.Channel, _ = strconv.Atoi(channel)
	friendAdd.RequestAt = time.Now().Format("2006-01-02 15:04:05")
	if err := global.Mysql.Create(&friendAdd).Error; err != nil {
		api.R(ctx, 500, "添加失败："+err.Error(), nil)
		return
	}

	api.Rt(ctx, 200, "成功发送添加请求", gin.H{})
}

// 删除好友
func DelFriend(ctx *gin.Context) {

}
