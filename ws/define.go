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
	Cmd    int    //指令
	FromId int    //来源id
	ToId   int    //接收id
	Ope    int    //消息通道
	Type   int    //消息类型
	Body   string //消息内容
}

// 本机连接表
var Connections = make(map[int]*Connection)

// 已认证连接表
var UidToClientId = make(map[int]string)

// 连接计数器&&并发锁
var ClientCounter = 0
var ClientCounterLocker sync.RWMutex

// 消息指令定义
const (
	CmdSuccess                = 1  //成功
	CmdFail                   = 2  //失败
	CmdSign                   = 3  //登录
	CmdSignSuccess            = 4  //登录成功
	CmdSignFail               = 5  //等失败
	CmdReceiveFriendMessage   = 6  //收到好友消息
	CmdWithdrawFriendMessage  = 7  //撤回好友消息
	CmdReceiveFriendAdd       = 8  //收到好友添加请求
	CmdReceiveFriendAddResult = 9  //收到好友请求结果
	CmdReceiveGroupMessage    = 10 //收到群消息
	CmdWithdrawGroupMessage   = 11 //撤回群消息
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
