package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ppIm/api"
	v1 "ppIm/api/v1"
	"ppIm/middleware"
	"ppIm/servers"
	"ppIm/ws"
)

func SetRouter(r *gin.Engine) {
	// 全局跨域中间件
	r.Use(middleware.Cors)

	// 公开访问目录
	r.StaticFS("/public", http.Dir("./public"))

	// websocket连接
	r.GET("/ws", ws.WebsocketEntry)
	// websocket服务状态
	r.GET("/ws/status", ws.StatusApi)
	//
	r.GET("/ws/isOnline", ws.IsOnlineApi)

	// 集群服务器列表
	r.GET("/cluster/servers", servers.ApiQuery)

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
		friend.POST("/add/confirm", v1.AddPass)
		// 删除好友
		friend.POST("/del", v1.DelFriend)

		// 群组系统
		group := im.Group("/group")
		// 创建群组
		group.POST("/create", v1.CreateGroup)
		// 群组列表
		group.POST("/list", v1.GroupList)
		// 请求加入群组
		group.POST("/join/request", v1.JoinGroup)
		// 加入群组请求处理
		group.POST("/join/confirm", v1.JoinPass)
		// 离开群组
		group.POST("/level", v1.LeaveGroup)
		// 设置群成员
		group.POST("/set/member", v1.SetMember)

		// 聊天系统
		chat := im.Group("/chat")
		// 发送消息给用户
		chat.POST("/send/user", v1.SendMessageToUser)
		// 撤回消息
		chat.POST("/withdraw/user", v1.WithdrawMessageFromUser)
	}

}
