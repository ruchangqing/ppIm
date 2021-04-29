package rpc

import (
	"golang.org/x/net/context"
	pb "ppIm/app/rpc/proto"
	"ppIm/app/websocket"
)

type imService struct{}

var ImService = imService{}

// 查询用户是否在线
func (h imService) IsOnline(ctx context.Context, in *pb.IsOnlineRequest) (*pb.IsOnlineResponse, error) {
	resp := new(pb.IsOnlineResponse)

	uid := int(in.Uid)
	resp.IsOnline = websocket.LocalIsOnline(uid)

	return resp, nil
}

// 发送给用户消息
func (h imService) SendToUser(ctx context.Context, in *pb.SendToUserRequest) (*pb.SendToUserResponse, error) {
	resp := new(pb.SendToUserResponse)
	resp.Result = false

	message := websocket.Message{
		Cmd:    int(in.Cmd),
		FromId: int(in.FromId),
		ToId:   int(in.ToId),
		Ope:    int(in.Ope),
		Type:   int(in.Type),
		Body:   in.Body,
	}
	websocket.LocalSendToUser(int(in.ToId), message)

	return resp, nil
}
