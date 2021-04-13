package rpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/rpc/proto" // 引入编译生成的包
)

func DialRpc() {
	// 连接
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Message)
}
