package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ppIm/app/http/api"
	"ppIm/app/model"
	"ppIm/lib"
	"ppIm/utils"
	"time"
)

type sign struct{}

var Sign sign

// 用户登录接口
func (sign) Login(ctx *gin.Context) {
	// 参数获取与验证
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	result, msg := validateParams(&username, &password)
	if false == result {
		api.R(ctx, api.Fail, msg, nil)
		return
	}

	// 检测用户是否存在
	var user model.User
	var count int
	lib.Db.Where("username = ?", username).Select("id,username,password,password_salt,nickname,avatar,status").First(&user).Count(&count)
	if count < 1 {
		api.R(ctx, api.Fail, "用户不存在，请更换用户名后重试", nil)
		return
	}

	// 验证密码是否合法
	postPassword := utils.Md5(utils.Md5(password) + user.PasswordSalt)
	if postPassword != user.Password {
		api.R(ctx, api.Fail, "密码错误", nil)
		return
	}

	// 密码正确，更新登陆时间
	loginTime := time.Now().Format("2006-01-02 15:04:05")
	lib.Db.Model(&user).Updates(map[string]interface{}{"login_time": loginTime, "last_ip": ctx.ClientIP()})

	// 生成jwt token和用户信息给用户
	tokenString := api.MakeJwtToken(user.Id)
	api.R(ctx, api.Success, "登录成功", gin.H{
		"t": tokenString,
		"user": gin.H{
			"username": username,
			"nickname": user.Nickname,
			"avatar":   utils.QiNiuClient.FullPath(user.Avatar),
			"status":   user.Status,
		},
	})
}

//用户注册接口
func (sign) Register(ctx *gin.Context) {
	// 参数获取与验证
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	result, msg := validateParams(&username, &password)
	if false == result {
		api.R(ctx, api.Fail, msg, nil)
		return
	}

	// 检测用户名是否存在
	var user model.User
	var count int
	lib.Db.Where("username = ?", username).First(&user).Count(&count)
	if count > 0 {
		api.R(ctx, api.Fail, "用户已存在，请更换用户名后重试", nil)
		return
	}

	// 新增用户数据，注册逻辑
	passwordSalt := utils.RandStr(6)
	password = utils.Md5(utils.Md5(password) + passwordSalt)

	user = model.User{
		Username:     username,
		Password:     password,
		Nickname:     "新用户" + username,
		PasswordSalt: passwordSalt,
		RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
		LoginTime:    time.Now().Format("2006-01-02 15:04:05"),
		Avatar:       viper.GetString("qiniu.default_avatar"),
		LastIp:       ctx.ClientIP(),
	}
	if err := lib.Db.Create(&user).Error; err != nil {
		lib.Logger.Debugf(err.Error())
		api.R(ctx, api.Fail, "服务器错误", nil)
		return
	}

	// 生成jwt token和用户信息给用户
	tokenString := api.MakeJwtToken(user.Id)
	api.R(ctx, api.Fail, "登录成功", gin.H{
		"t": tokenString,
		"user": gin.H{
			"username": username,
			"nickname": user.Nickname,
			"avatar":   utils.QiNiuClient.FullPath(user.Avatar),
			"status":   user.Status,
		},
	})
}

// --------------------- func --------------------- //

// 检测用户名、密码参数是否合法
func validateParams(username, password *string) (bool, string) {
	if len(*username) < 1 {
		return false, "请输入用户名"
	}
	if len(*username) < 6 || len(*username) > 20 {
		return false, "用户名长度限制6-20位"
	}
	if len(*password) < 1 {
		return false, "请输入密码"
	}
	if len(*password) < 6 || len(*password) > 20 {
		return false, "密码长度限制6-20位"
	}
	if utils.IsChinese(*username) || utils.IsChinese(*password) {
		return false, "用户名和密码不能含有中文"
	}
	return true, ""
}
