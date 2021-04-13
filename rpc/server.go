package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/rpc/proto" // 引入编译生成的包
)

const (
	// Address gRPC服务地址
	AddressS = "127.0.0.1:50052"
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

func main() {
	listen, err := net.Listen("tcp", AddressS)
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("Listen on " + AddressS)
	s.Serve(listen)
}
