package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"ppIm/app/rpc"
	pb "ppIm/app/rpc/proto"
	"ppIm/global"
	"ppIm/router"
	"ppIm/servers"
	"ppIm/utils"
	"time"
)

func StartServer() {
	go pprofServer()
	global.ServerAddress = utils.GetIntranetIp() + ":" + viper.GetString("cluster.rpc_port")
	go rpcServer()
	servers.Servers = servers.GetAllServers()
	servers.RegisterServer()
	go servers.WatchServers()
	httpServer()
}

// pprof性能监控服务，路由：/debug/pprof
func pprofServer() {
	log.Println(http.ListenAndServe(viper.GetString("http.pprof"), nil))
}

// rpc服务监听
func rpcServer() {
	listenAddress := "127.0.0.1:" + viper.GetString("cluster.rpc_port")
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic("开启RPC服务出错：" + err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterImServer(s, rpc.ImService)

	fmt.Println("[RPC-debug] Listen on " + listenAddress)
	fmt.Println(s.Serve(listen))
}

// http服务监听
func httpServer() {
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	gin.SetMode(viper.GetString("app.mode"))
	server := gin.Default()
	gin.DisableConsoleColor()
	// 日志使用文件
	path, _ := os.Getwd()
	file, err := os.Create(path + "/runtime/gin.log")
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
