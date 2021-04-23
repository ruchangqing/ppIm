package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"ppIm/api"
	"ppIm/global"
	"ppIm/model"
	"strconv"
	"time"
)

type group struct{}

var Group group

// 创建群组
func (group) Create(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	name := ctx.PostForm("name")
	if name == "" {
		api.R(ctx, global.FAIL, "请输入群名称", gin.H{})
		return
	}
	var group model.Group
	var count int
	global.Mysql.Where("name = ?", name).First(&group).Count(&count)
	if count > 0 {
		api.R(ctx, global.FAIL, "群组已存在", gin.H{})
		return
	}

	now := time.Now().Unix()
	trans := true
	// 事务
	global.Mysql.Transaction(func(tx *gorm.DB) error {
		group.OUid = uid
		group.Name = name
		group.CreatedAt = now
		if err := tx.Create(&group).Error; err != nil {
			trans = false
			return err
		}

		var groupUser model.GroupUser
		groupUser.GroupId = group.Id
		groupUser.UserId = uid
		groupUser.JoinAt = now
		if err := tx.Create(&groupUser).Error; err != nil {
			trans = false
			return err
		}

		return nil
	})

	if trans == true {
		api.Rt(ctx, global.SUCCESS, "创建成功", gin.H{})
	} else {
		api.R(ctx, global.FAIL, "创建失败", gin.H{})
	}
}

// 搜索群组
func (group) Search(ctx *gin.Context) {
	word := ctx.PostForm("word")
	if word == "" {
		api.R(ctx, global.FAIL, "请输入群组名称", gin.H{})
		return
	}
	var groups []model.Group
	global.Mysql.Model(&model.Group{}).Where("name LIKE ?", "%"+word+"%").Find(&groups)
	api.R(ctx, global.SUCCESS, "查询成功", gin.H{"list": groups})
}

// 我的群组
func (group) My(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	type Result struct {
		Id        int
		OUid      int
		Name      string
		CreatedAt int
		JoinAt    int
	}
	var result []Result
	global.Mysql.Raw("SELECT g.id,g.o_uid,g.name,g.created_at,u.join_at FROM `group_user` AS u INNER JOIN `group` AS g ON u.group_id = g.id WHERE user_id = ?", uid).Scan(&result)
	api.R(ctx, global.SUCCESS, "获取成功", gin.H{"list": result})
}

// 申请加入群组
func (group) Join(ctx *gin.Context) {
	uid := int(ctx.MustGet("id").(float64))
	groupId, _ := strconv.Atoi(ctx.PostForm("group_id"))
	var group model.Group
	global.Mysql.Where("id = ?", groupId).First(&group)
	if group.Id == 0 {
		api.R(ctx, global.FAIL, "群组不存在", gin.H{})
		return
	}
	if group.OUid == uid {
		api.R(ctx, global.FAIL, "你已经是群主", gin.H{})
		return
	}
	var groupUser model.GroupUser
	global.Mysql.Where("group_id = ? and user_id = ?", groupId, uid).First(&groupUser)
	if groupUser.Id > 0 {
		api.R(ctx, global.FAIL, "你已经在群组里", gin.H{})
		return
	}
	var groupJoin model.GroupJoin
	global.Mysql.Where("group_id = ? and user_id = ?", groupId, uid).First(&groupJoin)
	if groupJoin.Id > 0 {
		api.R(ctx, global.FAIL, "你的申请加入群组请求已经在处理中", gin.H{})
		return
	}
	groupJoin.GroupId = groupId
	groupJoin.UserId = uid
	groupJoin.JoinAt = time.Now().Unix()
	global.Mysql.Create(&groupJoin)
	if groupJoin.Id > 0 {
		api.Rt(ctx, global.SUCCESS, "申请成功", gin.H{})
	} else {
		api.R(ctx, global.FAIL, "申请失败", gin.H{})
	}
}

// 加入群组请求处理
func (group) JoinHandle(ctx *gin.Context) {

}

// 退出群组
func (group) Leave(ctx *gin.Context) {

}

// 设置成员
func (group) SetMember(ctx *gin.Context) {

}
