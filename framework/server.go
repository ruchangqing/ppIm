package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ppIm/router"
	"ppIm/rpc"
	"time"
)

func StartServer() {
	// 测试rpc
	go rpc.Server()
	go func() {
		timer := time.NewTicker(3 * time.Second)
		for  {
			<-timer.C
			rpc.DialRpc()
		}
	}()

	httpServer()
}

func httpServer() {
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	gin.SetMode(viper.GetString("app.mode"))
	server := gin.Default()
	/*	gin.DisableConsoleColor()
			// 日志使用文件
			file, _ := os.Create("./runtime/access.log")
			gin.DefaultWriter = file
		/*	// 自定义日志格式
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
	*/
	server.Use(gin.Recovery())
	router.SetRouter(server)
	panic(server.Run(fmt.Sprintf("%s:%s", host, port)))
}
