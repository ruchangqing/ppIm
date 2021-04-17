package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/servers/rpc/proto" // 引入编译生成的包
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func Server() {
	rpcAddress := "127.0.0.1:" + viper.GetString("cluster.rpc_port")
	listen, err := net.Listen("tcp", rpcAddress)
	if err != nil {
		fmt.Println(err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("[RPC-debug] Listen on " + rpcAddress)
	fmt.Println(s.Serve(listen))
}
