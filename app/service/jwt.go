package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"ppIm/lib"
	"time"
)

// 解析token
func ParseToken(ctx *gin.Context, jwtToken string) (int, string) {
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return lib.JwtHmacSampleSecret, nil
	})
	// 开始解析token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 解析成功
		ctx.Set("id", claims["id"])
		ctx.Set("at", claims["at"])
		at := int64(ctx.MustGet("at").(float64)) // token生成时间戳
		id := int(ctx.MustGet("id").(float64))
		nowAt := time.Now().Unix() // 当前时间戳
		expireAt := at + 7*24*3600 // 过期时间戳，一周后
		if nowAt > expireAt {
			return 0, "登录状态已过期，请重新登录"
		} else {
			return id, ""
		}
	} else {
		return 0, "鉴权失败"
	}
}
