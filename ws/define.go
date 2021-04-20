package ws

import (
	"github.com/gorilla/websocket"
	"ppIm/global"
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
	MsgType    int
	MsgContent string
}

// 本机连接表
var Connections = make(map[int]*Connection)

// 已认证连接表
var UidToClientId = make(map[int]string)

// 连接计数器&&并发锁
var ClientCounter = 0
var ClientCounterLocker sync.RWMutex

// 消息类型定义
const (
	OK                     = 90001 //成功
	Fail                   = 90002 //失败
	Sign                   = 10000 //登录
	SignSuccess            = 10001 //登录成功
	SignFail               = 10002 //等失败
	ReceiveFriendMessage   = 20001 //收到好友消息
	RecallFriendMessage    = 20002 //撤回私聊消息
	ReceiveFriendAdd       = 30001 //收到好友添加请求
	ReceiveFriendAddResult = 30002 //收到好友请求结果
)

//生成clientId
func GenClientId(clientCounter int) string {
	str := global.ServerAddress + "@@" + strconv.Itoa(clientCounter)
	//clientId, err := utils.AesEncrypt([]byte(str))
	//if err != nil {
	//	fmt.Println("生成clientId出错：" + err.Error())
	//}
	clientId := str
	return clientId
}
