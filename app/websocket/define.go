package websocket

import (
	"github.com/gorilla/websocket"
	"ppIm/lib"
	"strconv"
	"sync"
)

// 连接结构体
type Connection struct {
	ClientId string
	Uid      int
	Conn     *websocket.Conn
}

// 消息结构体
type Message struct {
	Cmd    int    //指令
	FromId int    //来源id
	ToId   int    //接收id
	Ope    int    //消息通道
	Type   int    //消息类型
	Body   string //消息内容
}

// 消息指令定义
const (
	CmdSuccess                = 1  //成功
	CmdFail                   = 2  //失败
	CmdSign                   = 3  //登录
	CmdSignSuccess            = 4  //登录成功
	CmdReceiveFriendMessage   = 6  //收到好友消息
	CmdWithdrawFriendMessage  = 7  //撤回好友消息
	CmdReceiveFriendAdd       = 8  //收到好友添加请求
	CmdReceiveFriendAddResult = 9  //收到好友请求结果
	CmdReceiveGroupMessage    = 10 //收到群消息
	CmdWithdrawGroupMessage   = 11 //撤回群消息
	CmdReceiveGroupJoin       = 12 //收到加入群组请求
	CmdReceiveGroupJoinResult = 13 //收到加入群组结果
	CmdReceiveGroupShot       = 14 //收到被踢出群组通知
)

// 本机连接表
var Connections = make(map[int]*Connection)
// 连接计数器&&并发锁
var ClientCounter = 0
var ClientCounterLocker sync.RWMutex

// 已认证连接表
var UidToClientId = make(map[int]string)
var UidToClientIdLocker sync.RWMutex


//生成clientId
func GenClientId(clientCounter int) string {
	str := lib.ServerAddress + "@@" + strconv.Itoa(clientCounter)
	clientId := str
	return clientId
}
