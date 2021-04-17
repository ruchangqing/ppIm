package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"ppIm/global"
	"ppIm/router"
	"ppIm/servers"
	"ppIm/servers/rpc"
	"time"
)

func StartServer() {
	if global.IsCluster {
		go rpc.Server()
		servers.Servers = servers.GetAllServers()
		servers.RegisterServer()
		go servers.WatchServers()
	}
	httpServer()
}

func httpServer() {
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	gin.SetMode(viper.GetString("app.mode"))
	server := gin.Default()
	gin.DisableConsoleColor()
	// 日志使用文件
	path, _ := os.Getwd()
	file, err := os.Create(path + "/runtime/access.log")
	if err != nil {
		fmt.Println(err)
	}
	gin.DefaultWriter = file
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

	server.Use(gin.Recovery())
	router.SetRouter(server)
	listenAddress := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("[GIN-debug] Listen on " + listenAddress)
	fmt.Println(server.Run(listenAddress))
}
