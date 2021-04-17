package ws

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ppIm/servers/rpc/proto"
)

// RPC远程调用查询是否在线
func RpcIsOnline(rpcAddress string, uid int) bool {
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接rpc服务器" + rpcAddress + " 发生错误：" + err.Error())
		return false
	}
	defer conn.Close()

	c := pb.NewImClient(conn)

	req := &pb.IsOnlineRequest{Uid: int64(uid)}
	res, err := c.IsOnline(context.Background(), req)

	if err != nil {
		fmt.Println("调用rpc服务器 " + rpcAddress + " IsOnline方法发生错误：" + err.Error())
		return false
	}

	return res.IsOnline
}

// RPC远程调用发送给用户消息
func RpcSendToUser(rpcAddress string, targetUid int, msgType int, msgContent string) bool {
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接rpc服务器" + rpcAddress + " 发生错误：" + err.Error())
		return false
	}
	defer conn.Close()

	c := pb.NewImClient(conn)

	req := &pb.SendToUserRequest{
		TargetUid:  int64(targetUid),
		MsgType:    int64(msgType),
		MsgContent: msgContent,
	}
	res, err := c.SendToUser(context.Background(), req)

	if err != nil {
		fmt.Println("调用rpc服务器 " + rpcAddress + " IsOnline方法发生错误：" + err.Error())
		return false
	}

	return res.Result
}
