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

// 好友列表
func FriendList(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	type Result struct {
		Id       int
		Nickname string
		Username string
		Avatar   string
		Sex      int
	}
	var result Result
	var results []Result
	rows, err := global.Mysql.Raw("select u.id,u.nickname,u.username,u.avatar,u.sex from friend_list as f join user as u on f.f_uid = u.id where f.uid = ?", uid).Rows()
	if err != nil {
		print(err)
		api.R(ctx, 500, "未知错误", nil)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.Id, &result.Nickname, &result.Username, &result.Avatar, &result.Sex)
		results = append(results, result)
	}
	api.Rt(ctx, 200, "请求成功", gin.H{"lists": results})
}

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
	friendAdd.Channel = channel
	friendAdd.RequestAt = time.Now().Format("2006-01-02 15:04:05")
	if err := global.Mysql.Create(&friendAdd).Error; err != nil {
		api.R(ctx, 500, "添加失败："+err.Error(), nil)
		return
	}

	// 实时通知用户添加请求
	if ws.IsOnline(user.Id) {
		var me model.User
		global.Mysql.Where("id = ?", uid).First(&me)
		ws.Connections[user.Id].Conn.WriteJSON(ws.WsMsg(3, 200, me.Nickname+"请求添加您为好友", gin.H{
			"uid":      uid,
			"nickname": me.Nickname,
			"avatar":   me.Avatar,
			"username": me.Username,
			"reason":   reason,
			"channel":  channel,
		}))
	}

	api.Rt(ctx, 200, "成功发送添加请求", gin.H{})
}

// 收到的好友请求列表
func AddList(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	type Result struct {
		Id        int
		Nickname  string
		Username  string
		Avatar    string
		Channel   string
		Reason    string
		RequestAt string
	}
	var result Result
	var results []Result
	rows, err := global.Mysql.Raw("select u.id,u.nickname,u.username,u.avatar,f.channel,f.reason,f.request_at from friend_add as f join user as u on f.uid = u.id  where f.f_uid = ? and f.status = 0 order by request_at desc", uid).Rows()
	if err != nil {
		print(err)
		api.R(ctx, 500, "未知错误", nil)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.Id, &result.Nickname, &result.Username, &result.Avatar, &result.Channel, &result.Reason, &result.RequestAt)
		results = append(results, result)
	}
	api.Rt(ctx, 200, "请求成功", gin.H{"list": results})
}

// 处理收到的好友请求
func AddPass(ctx *gin.Context) {
	// 添加好友请求验证与数据写入
	fUid, _ := strconv.Atoi(ctx.PostForm("f_uid"))
	status, _ := strconv.Atoi(ctx.PostForm("status"))
	if status != 1 && status != -1 {
		api.R(ctx, 500, "非法参数", nil)
		return
	}
	uid := int(ctx.MustGet("id").(float64))
	var friendAdd model.FriendAdd
	global.Mysql.Where("uid = ? and f_uid = ? and status = ?", fUid, uid, 0).First(&friendAdd)
	if friendAdd.Id == 0 {
		api.R(ctx, 500, "添加好友请求不存在", nil)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	global.Mysql.Model(&friendAdd).Updates(map[string]interface{}{"status": status, "pass_at": now})

	// 通过验证后进行好友数据写入
	friendList1 := &model.FriendList{
		Uid:       uid,
		FUid:      fUid,
		Channel:   friendAdd.Channel,
		Reason:    friendAdd.Reason,
		Role:      1,
		CreatedAt: now,
	}
	friendList2 := &model.FriendList{
		Uid:       fUid,
		FUid:      uid,
		Channel:   friendAdd.Channel,
		Reason:    friendAdd.Reason,
		Role:      2,
		CreatedAt: now,
	}
	err1 := global.Mysql.Create(&friendList1).Error
	err2 := global.Mysql.Create(&friendList2).Error
	if err1 != nil || err2 != nil {
		api.R(ctx, 500, "未知错误", nil)
		return
	}

	// 实时通知用户我是否同意
	if ws.IsOnline(fUid) {
		var me model.User
		global.Mysql.Where("id = ?", uid).First(&me)
		var msg string
		if status == 1 {
			msg = "同意您添加为好友"
		} else if status == -1 {
			msg = "拒绝您添加为好友"
		}
		ws.Connections[fUid].Conn.WriteJSON(ws.WsMsg(4, 200, me.Nickname+msg, gin.H{
			"uid":         uid,
			"pass_status": status,
			"nickname":    me.Nickname,
			"avatar":      me.Avatar,
			"username":    me.Username,
		}))
	}

	api.Rt(ctx, 200, "成功", gin.H{})
}

// 删除好友
func DelFriend(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	fUid, _ := strconv.Atoi(ctx.PostForm("f_uid"))
	var friendList model.FriendList
	global.Mysql.Where("uid = ? and f_uid = ?", uid, fUid).First(&friendList)
	if friendList.Id == 0 {
		api.R(ctx, 500, "对方不是你的好友", nil)
		return
	}
	err := global.Mysql.Delete(&friendList).Error
	if err != nil {
		print(err)
		api.R(ctx, 500, "删除失败", nil)
		return
	}
	api.Rt(ctx, 200, "删除成功", gin.H{})
}
