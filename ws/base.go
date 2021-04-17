package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
	"ppIm/middleware"
	"ppIm/utils"
	"strconv"
	"sync"
	"time"
)

// 连接结构体
type Connection struct {
	ClientId string
	Uid      int
	Conn     *websocket.Conn
}

// 本机连接表
var Connections = make(map[int]*Connection)

// 已认证连接表
var UidToClientId = make(map[int]string)

// 连接计数器&&并发锁
var ClientCounter = 0
var ClientCounterLocker sync.RWMutex

//生成clientId
func GenClientId(clientCounter int) string {
	serverAddress := utils.GetIntranetIp() + ":" + viper.GetString("cluster.rpc_port")
	str := serverAddress + "@@" + strconv.Itoa(clientCounter)
	//clientId, err := utils.AesEncrypt([]byte(str))
	//if err != nil {
	//	fmt.Println("生成clientId出错：" + err.Error())
	//}
	clientId := str
	return clientId
}

// 升级http为websocket服务
var WebsocketUpgrade = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocket服务入口
func WebsocketEntry(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	var conn *websocket.Conn
	var err error
	conn, err = WebsocketUpgrade.Upgrade(w, r, nil) // 升级为websocket协议
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	var c Connection
	var counter int

	ClientCounterLocker.Lock()
	ClientCounter++
	counter = ClientCounter
	Connections[counter] = &c
	ClientCounterLocker.Unlock()

	c.ClientId = GenClientId(counter)

	// 15秒内没收到token绑定成功的断开连接
	time.AfterFunc(15*time.Second, func() {
		if c.Uid == 0 {
			conn.Close()
		}
	})

	// 必须死循环，gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// json消息转换
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			// 接收到非json数据
			fmt.Println("[Websocket]Message json.Unmarshal fail: " + string(msg))
			conn.WriteJSON(WsMsg(-1, 0, "非json格式数据", nil))
			continue
		}

		// bind绑定uid和client_id，这是必须绑定的才能通信的
		if message.Cmd == 1 {
			if c.Uid == 0 {
				jwtToken := message.Data["token"]
				id, err := middleware.ParseToken(ctx, jwtToken.(string))
				if err != "" {
					fmt.Println(err)
				}
				c.Uid = id
				UidToClientId[c.Uid] = c.ClientId // 认证成功后注册到已认证连接表，方便查询对应clientId
				conn.WriteJSON(WsMsg(1, 1, "ok", nil))
			}
		} else {
			if c.Uid > 0 {
				Receive(&c, message)
			} else {
				conn.WriteJSON(WsMsg(-1, 0, "Not bind!", nil))
			}
		}
	}

	ClientCounterLocker.Lock()
	delete(Connections, counter)
	if c.Uid > 0 {
		delete(UidToClientId, c.Uid)
	}
	ClientCounterLocker.Unlock()
}
