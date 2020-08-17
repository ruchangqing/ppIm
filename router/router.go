package router

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	v1 "ppIm/api/v1"
	"ppIm/middleware"
	"ppIm/ws"
)

func SetRouter(r *gin.Engine) {

	r.GET("/ws", ws.WsHandler)

	// 首页
	r.GET("/", api.Welcome)
	// 未定义路由
	r.NoRoute(api.NotFound)
	r.NoMethod(api.NotFound)

	// 用户登录
	r.POST("/api/v1/login", v1.Login)
	// 用户注册
	r.POST("/api/v1/register", v1.Register)

	// 用户相关接口
	user := r.Group("/api/v1/user")
	user.Use(middleware.ValidateJwtToken)
	{
		// 设置昵称
		user.POST("/update/nickname", v1.UpdateNickname)
		// 设置头像
		user.POST("/update/avatar", v1.UpdateAvatar)
		// 实名认证
		user.POST("/update/realname", v1.RealNameVerify)
		// 更新用户位置（经纬度）
		user.POST("/update/location", v1.UpdateLocation)
	}

	// 位置相关接口
	geo := r.Group("/api/v1/geo")
	geo.Use(middleware.ValidateJwtToken)
	{
		// 用户列表（按距离排序）
		geo.POST("/users", v1.GeoUsers)
	}

}
