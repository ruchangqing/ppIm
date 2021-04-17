package im

import (
	"fmt"
	"net"
	"ppIm/global"
	pb "ppIm/servers/rpc/im/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type imService struct{}

var ImService = imService{}

// 查询用户是否在线
func (h imService) IsOnline(ctx context.Context, in *pb.IsOnlineRequest) (*pb.IsOnlineResponse, error) {
	resp := new(pb.IsOnlineResponse)
	resp.IsOnline = false

	return resp, nil
}

// 发送给用户消息
func (h imService) SendToUser(ctx context.Context, in *pb.SendToUserRequest) (*pb.SendToUserResponse, error) {
	resp := new(pb.SendToUserResponse)
	resp.Result = false

	return resp, nil
}

// 启动RPC服务
func Server() {
	listen, err := net.Listen("tcp", global.ServerAddress)
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	pb.RegisterImServer(s, ImService)

	fmt.Println("[RPC-debug] Listen on " + global.ServerAddress)
	fmt.Println(s.Serve(listen))
}
