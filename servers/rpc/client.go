package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/servers/rpc/proto" // 引入编译生成的包
)

func DialRpc() {
	rpcAddress := viper.GetString("cluster.rpc_address")
	// 连接
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
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
