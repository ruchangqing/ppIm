package v1

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	"ppIm/global"
	"ppIm/model"
	"time"
)

type group struct{}

var Group group

// 创建群组
func (group) Create(ctx *gin.Context) {
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
	global.Mysql.Create(&model.Group{Name: name, CreatedAt: time.Now().Unix()})
	api.Rt(ctx, global.SUCCESS, "创建成功", gin.H{})
}

// 群组列表
func (group) List(ctx *gin.Context) {

}

// 我的群组
func (group) My(ctx *gin.Context) {

}

// 加入群组
func (group) Join(ctx *gin.Context) {

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
