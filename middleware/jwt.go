package middleware

import (
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

	// 首先redis查询token是否有效
	cacheKey := fmt.Sprintf("user:token:%s", jwtToken)
	_, err := global.Redis.Get(global.RedisCtx, cacheKey).Result()
	if err == redis.Nil {
		api.R(ctx, 401, "鉴权失败:-2", nil)
		ctx.Abort()
		return
	}

	// 开始解析token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return global.JwtHmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 解析成功
		ctx.Set("id", claims["id"])
		ctx.Set("at", claims["at"])
		at := int64(ctx.MustGet("at").(float64)) // token生成时间戳
		nowAt := time.Now().Unix()               // 当前时间戳
		expireAt := at + 7*24*3600               // 过期时间戳，一周后
		if nowAt > expireAt {
			api.R(ctx, 401, "登陆状态已过期", nil)
			ctx.Abort()
			return
		}
	} else {
		// 解析失败，响应结束
		fmt.Println(err)
		api.R(ctx, 403, "鉴权失败:-3", nil)
		ctx.Abort()
		return
	}
}
