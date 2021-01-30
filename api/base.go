package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"ppIm/global"
	"time"
)

func Welcome(ctx *gin.Context) {
	R(ctx, global.SUCCESS, "欢迎使用ppim开源社交项目，定制开发联系微信：liboy825", nil)
}

func NotFound(ctx *gin.Context) {
	R(ctx, global.NOTFOUND, "Not Found", nil)
}

// 响应封装（不带新token）
func R(ctx *gin.Context, code int, msg string, data gin.H) {
	responseType := viper.GetString("http.responseType")
	data = gin.H{"code": code, "msg": msg, "data": data}
	switch responseType {
	case "json":
		ctx.JSON(http.StatusOK, data)
	case "xml":
		ctx.XML(http.StatusOK, data)
	default:
		ctx.JSON(http.StatusOK, data)
	}
}

// 响应封装（带新token）
func Rt(ctx *gin.Context, code int, msg string, data gin.H) {
	/*	// 为实现7天内活跃用户无需重新登陆，且生成的的token只能使用一次：请求成功响应后删掉老token，生成新token
		jwtToken := ctx.GetHeader("Login-Token")
		cacheKey := fmt.Sprintf("user:token:%s", jwtToken)
	*/

	// 为实现7天内活跃用户无需重新登陆：请求成功响应后生成新token
	id := int(ctx.MustGet("id").(float64))
	//global.Redis.Del(context.Background(), cacheKey)
	data["t"] = MakeJwtToken(id)

	responseType := viper.GetString("http.responseType")
	data = gin.H{"code": code, "msg": msg, "data": data}
	switch responseType {
	case "json":
		ctx.JSON(http.StatusOK, data)
	case "xml":
		ctx.XML(http.StatusOK, data)
	default:
		ctx.JSON(http.StatusOK, data)
	}
}

// 生成用户token
func MakeJwtToken(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"at": time.Now().Unix(), // token有效期：一天
	})
	tokenString, err := token.SignedString(global.JwtHmacSampleSecret)
	if err != nil {
		print(err)
	}

	// 缓存用户最新token到redis，实现token只能使用一次
	//cacheKey := fmt.Sprintf("user:token:%s", tokenString)
	//global.Redis.Set(context.Background(), cacheKey, id, 7*24*3600*time.Second)

	return tokenString
}
