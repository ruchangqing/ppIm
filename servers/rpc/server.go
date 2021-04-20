package rpc

import (
	"golang.org/x/net/context"
	pb "ppIm/servers/rpc/proto"
	"ppIm/ws"
	"strconv"
	"strings"
)

type imService struct{}

var ImService = imService{}

// 查询用户是否在线
func (h imService) IsOnline(ctx context.Context, in *pb.IsOnlineRequest) (*pb.IsOnlineResponse, error) {
	resp := new(pb.IsOnlineResponse)

	uid := int(in.Uid)
	resp.IsOnline = ws.IsOnline(uid)

	return resp, nil
}

// 发送给用户消息
func (h imService) SendToUser(ctx context.Context, in *pb.SendToUserRequest) (*pb.SendToUserResponse, error) {
	resp := new(pb.SendToUserResponse)
	resp.Result = false

	clientId := ws.UidToClientId[int(in.TargetUid)]
	arr := strings.Split(clientId, "@@")
	number, _ := strconv.Atoi(arr[1])
	message := ws.Message{int(in.MsgType), in.MsgContent}
	ws.Connections[number].Conn.WriteJSON(message)

	return resp, nil
}
