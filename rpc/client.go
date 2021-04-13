package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/rpc/proto" // 引入编译生成的包
)

const (
	// Address gRPC服务地址
	AddressC = "127.0.0.1:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(AddressC, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Message)
}
