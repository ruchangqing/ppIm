package middleware

import (
	"github.com/gin-gonic/gin"
	"ppIm/app/http/api"
	"ppIm/app/service"
)

// 检查Token是否有效
func ValidateJwtToken(ctx *gin.Context) {
	// 校验header中token值格式是否合法
	jwtToken := ctx.GetHeader("Login-Token")
	if len(jwtToken) < 16 {
		api.R(ctx, api.NoAuth, "鉴权失败", nil)
		ctx.Abort()
		return
	}

	_, err := service.ParseToken(ctx, jwtToken)
	if err != "" {
		// 解析失败，响应结束
		api.R(ctx, api.NoAuth, err, nil)
		ctx.Abort()
		return
	}
}
