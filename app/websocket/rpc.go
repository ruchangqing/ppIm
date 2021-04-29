package websocket

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/app/rpc/proto"
	"ppIm/lib"
)

// RPC远程调用查询是否在线
func RpcIsOnline(rpcAddress string, uid int) bool {
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		lib.Logger.Debugf("连接rpc服务器" + rpcAddress + " 发生错误：" + err.Error())
		return false
	}
	defer conn.Close()

	c := pb.NewImClient(conn)

	req := &pb.IsOnlineRequest{Uid: int64(uid)}
	res, err := c.IsOnline(context.Background(), req)

	if err != nil {
		lib.Logger.Debugf("调用rpc服务器 " + rpcAddress + " IsOnline方法发生错误：" + err.Error())
		return false
	}

	return res.IsOnline
}

// RPC远程调用发送给用户消息
func RpcSendToUser(rpcAddress string, message Message) bool {
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		lib.Logger.Debugf("连接rpc服务器" + rpcAddress + " 发生错误：" + err.Error())
		return false
	}
	defer conn.Close()

	c := pb.NewImClient(conn)

	req := &pb.SendToUserRequest{
		Cmd:    int64(message.Cmd),
		FromId: int64(message.FromId),
		ToId:   int64(message.ToId),
		Ope:    int64(message.Ope),
		Type:   int64(message.Type),
		Body:   message.Body,
	}
	res, err := c.SendToUser(context.Background(), req)

	if err != nil {
		lib.Logger.Debugf("调用rpc服务器 " + rpcAddress + " SendToUser方法发生错误：" + err.Error())
		return false
	}

	return res.Result
}
