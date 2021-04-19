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
	Sign                 = 10000
	SignSuccess          = 10001
	SignFail             = 10002
	Fail                 = 90000
	ReceiveFriendMessage = 20001
	RecallFriendMessage  = 20002
	ReceiveFriendAdd     = 30001
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
