package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"ppIm/api"
	"ppIm/global"
	"time"
)

// 鉴权token
func ValidateJwtToken(ctx *gin.Context) {
	// 校验header中token值格式是否合法
	jwtToken := ctx.GetHeader("Login-Token")
	if len(jwtToken) < 16 {
		api.R(ctx, 401, "鉴权失败:-1", nil)
		ctx.Abort()
		return
	}

	_, err := ParseToken(ctx, jwtToken)
	if err != "" {
		// 解析失败，响应结束
		api.R(ctx, 403, err, nil)
		ctx.Abort()
		return
	}

	// redis查询token是否有效
	cacheKey := fmt.Sprintf("user:token:%s", jwtToken)
	_, err2 := global.Redis.Get(context.Background(), cacheKey).Result()
	if err2 == redis.Nil {
		api.R(ctx, 401, "鉴权失败:-2", nil)
		ctx.Abort()
		return
	}

}

// 解析token
func ParseToken(ctx *gin.Context, jwtToken string) (int, string) {
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return global.JwtHmacSampleSecret, nil
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
		return 0, "鉴权失败:-3"
	}
}
