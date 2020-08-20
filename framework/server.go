package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"ppIm/router"
	"time"
)

func StartServer() {
	httpServer()
}

func httpServer() {
	gin.SetMode(viper.GetString("app.mode"))
	gin.DisableConsoleColor()
	// 日志使用文件
	file, _ := os.Create("./runtime/access.log")
	gin.DefaultWriter = file
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	server := gin.Default()
	// 自定义日志格式
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
	}))
	// 防止日志中间件报错终止程序运行
	server.Use(gin.Recovery())
	router.SetRouter(server)
	panic(server.Run(fmt.Sprintf("%s:%s", host, port)))
}
