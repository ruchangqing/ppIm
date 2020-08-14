package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ppIm/global"
	"time"
)

func Welcome(ctx *gin.Context) {
	R(ctx, 200, "Thank you for use PPF.", nil)
}

func NotFound(ctx *gin.Context) {
	R(ctx, 404, "Not Found", nil)
}

// 响应封装（不带新token）
func R(ctx *gin.Context, code int, msg string, data gin.H) {
	responseType := viper.GetString("http.responseType")
	data = gin.H{"code": code, "msg": msg, "data": data}
	switch responseType {
	case "json":
		ctx.JSON(code, data)
	case "xml":
		ctx.XML(code, data)
	default:
		ctx.JSON(code, data)
	}
}

// 响应封装（带新token）
func Rt(ctx *gin.Context, code int, msg string, data gin.H) {
	// 为实现7天内活跃用户无需重新登陆，且生成的的token只能使用一次：请求成功响应后删掉老token，生成新token
	jwtToken := ctx.GetHeader("Login-Token")
	cacheKey := fmt.Sprintf("user:token:%s", jwtToken)
	id := uint(ctx.MustGet("id").(float64))
	global.Redis.Del(global.RedisCtx, cacheKey)
	data["t"] = MakeJwtToken(id)

	responseType := viper.GetString("http.responseType")
	data = gin.H{"code": code, "msg": msg, "data": data}
	switch responseType {
	case "json":
		ctx.JSON(code, data)
	case "xml":
		ctx.XML(code, data)
	default:
		ctx.JSON(code, data)
	}
}

// 生成用户token
func MakeJwtToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"at": time.Now().Unix(), // token有效期：一天
	})
	tokenString, err := token.SignedString(global.JwtHmacSampleSecret)
	if err != nil {
		print(err)
	}

	// 缓存用户最新token到redis，实现token只能使用一次
	cacheKey := fmt.Sprintf("user:token:%s", tokenString)
	global.Redis.Set(global.RedisCtx, cacheKey, id, 7*24*3600*time.Second)

	return tokenString
}
