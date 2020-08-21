package router

import (
	"github.com/gin-gonic/gin"
	"ppIm/api"
	v1 "ppIm/api/v1"
	"ppIm/middleware"
	"ppIm/ws"
)

func SetRouter(r *gin.Engine) {
	// 全局跨域中间件
	r.Use(middleware.Cors)

	// 测试接口
	r.GET("/test", v1.Test)

	// websocket连接
	r.GET("/ws", ws.WebsocketEntry)
	// websocket服务状态
	r.GET("/ws/status", ws.Status)

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
		// 附近的人
		geo.POST("/users", v1.GeoUsers)
	}

	// 聊天相关接口
	im := r.Group("/api/v1/im")
	im.Use(middleware.ValidateJwtToken)
	{

		// 好友系统
		friend := im.Group("/friend")
		//好友列表
		friend.POST("/list", v1.FriendList)
		// 添加好友
		friend.POST("/add/request", v1.AddFriend)
		// 收到的添加请求列表
		friend.POST("/add/list", v1.AddList)
		// 处理好友请求
		friend.POST("/add/confirm", v1.AddConfirm)
		// 删除好友
		friend.POST("/del", v1.DelFriend)

	}

}
