package servers

import (
	"github.com/gin-gonic/gin"
	"ppIm/app/http/api"
	"sync"
)

var Servers []string
var ServersLock sync.RWMutex

//添加服务器
func AddServer(server string) {
	ServersLock.Lock()
	exists := false
	for i := 0; i < len(Servers); i++ {
		if Servers[i] == server {
			exists = true
			break
		}
	}
	if !exists {
		Servers = append(Servers, server)
	}
	ServersLock.Unlock()
}

//删除服务器
func DelServer(server string) {
	ServersLock.Lock()
	for i := 0; i < len(Servers); i++ {
		if Servers[i] == server {
			Servers = append(Servers[:i], Servers[i+1:]...)
			i--
		}
	}
	ServersLock.Unlock()
}

//查询服务器接口
func Api(ctx *gin.Context) {
	api.R(ctx, 200, "status", gin.H{
		"servers": Servers,
	})
}