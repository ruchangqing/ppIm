package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"ppIm/app/api"
	"ppIm/app/model"
	"ppIm/global"
	"ppIm/utils"
	"strconv"
	"strings"
	"time"
)

type user struct{}

var User user

// 用户更新昵称
func (user) UpdateNickname(ctx *gin.Context) {
	// 校验昵称参数
	nickname := ctx.PostForm("nickname")
	if len(nickname) < 4 || len(nickname) > 20 {
		api.R(ctx, global.ApiFail, "昵称长度不合法", nil)
		return
	}

	// jwt参数
	id := int(ctx.MustGet("id").(float64))
	// 更新用户昵称
	user := &model.User{Id: id}
	result := global.Db.Model(&user).Update("nickname", nickname).RowsAffected
	api.Rt(ctx, global.ApiSuccess, "设置成功", gin.H{"result": result})
}

// 用户更新头像
func (user) UpdateAvatar(ctx *gin.Context) {

	file, err := ctx.FormFile("avatar")
	if err != nil {
		api.R(ctx, global.ApiFail, "请选择图片", nil)
		return
	}
	if file.Size/1024 > 2048 {
		api.R(ctx, global.ApiFail, "图片大小限制在2mb", nil)
		return
	}

	// 保存头像文件，格式为id
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".bmp" && fileExt != ".png" && fileExt != ".gif" && fileExt != ".tif" {
		api.R(ctx, global.ApiFail, "图片格式不受支持", nil)
		return
	}
	id := int(ctx.MustGet("id").(float64))
	now := time.Now().Unix()
	// 本地缓存地址
	localPath := fmt.Sprintf("runtime/upload/%d_%d%s", id, now, fileExt)
	if err := ctx.SaveUploadedFile(file, localPath); err != nil {
		api.R(ctx, global.ApiFail, "服务器错误", nil)
		global.Logger.Debugf(err.Error())
		return
	}
	// 七牛云上传地址
	uploadPath := fmt.Sprintf("avatar/%d_%d%s", id, now, fileExt)
	err = utils.QiNiuClient.Upload(localPath, uploadPath)
	// 删除本地缓存
	os.Remove(localPath)
	if err != nil {
		api.R(ctx, global.ApiFail, "服务器错误", nil)
		global.Logger.Debugf(err.Error())
		return
	}

	// 更新头像地址到数据库
	user := &model.User{Id: id}
	user.Avatar = uploadPath
	result := global.Db.Model(&user).Update(user).RowsAffected

	api.Rt(ctx, global.ApiSuccess, "设置成功", gin.H{"result": result})
}

// 实名认证
func (user) RealNameVerify(ctx *gin.Context) {
	// 校验参数
	realName := ctx.PostForm("real_name")
	idCard := ctx.PostForm("id_card")
	if len(realName) < 4 || len(realName) > 20 || !utils.IsChinese(realName) {
		api.R(ctx, global.ApiFail, "姓名长度不合法", nil)
		return
	}

	//校验身份证信息
	x := []byte(idCard)
	if validate := utils.IsValidCitizenNo(&x); !validate {
		api.R(ctx, global.ApiFail, "身份证不合法", nil)
		return
	}
	// 获取身份证信息：性别、生日、省份
	_, _, sex, _ := utils.GetCitizenNoInfo(x)
	uSex := 0
	if sex == "男" {
		uSex = 1
	} else if sex == "女" {
		uSex = 2
	}

	// jwt参数
	id := int(ctx.MustGet("id").(float64))
	// 更新实名信息
	user := &model.User{Id: id}
	result := global.Db.Model(&user).Updates(map[string]interface{}{"real_name": realName, "id_card": idCard, "sex": uSex}).RowsAffected
	api.Rt(ctx, global.ApiSuccess, "实名认证成功", gin.H{"result": result})
}

//  更新最新地理位置及IP
func (user) UpdateLocation(ctx *gin.Context) {
	longitude := ctx.PostForm("longitude")
	latitude := ctx.PostForm("latitude")

	/*
	   isLong := regexp.MustCompile("^(([1-9]\\\\d?)|(1[0-7]\\\\d))(\\\\.\\\\d{1,6})|180|0(\\\\.\\\\d{1,6})?$")
	   	isLatitude := regexp.MustCompile("^(([1-8]\\\\d?)|([1-8]\\\\d))(\\\\.\\\\d{1,6})|90|0(\\\\.\\\\d{1,6})?$")
	   	if !isLong.MatchString(longitude) || !isLatitude.MatchString(latitude) {
	   		api.R(ctx, 500, "位置信息不合法", nil)
	   		return
	   	}
	*/

	id := int(ctx.MustGet("id").(float64))
	user := &model.User{Id: id}
	result := global.Db.Model(&user).Updates(map[string]interface{}{"longitude": longitude, "latitude": latitude, "last_ip": ctx.ClientIP()}).RowsAffected
	api.Rt(ctx, global.ApiSuccess, "更新位置成功", gin.H{"result": result})

	// 更新经纬度到es，用于后期查询
	data := fmt.Sprintf(`{
    "uid": "%d",
    "location": "%s,%s"
    }`, id, latitude, longitude)
	_, err := global.Elasticsearch.Index().Index("user_location").Id(strconv.Itoa(int(id))).BodyJson(data).Do(context.Background())
	if err != nil {
		global.Logger.Debugf(err.Error())
	}
}
