package v1

import (
	"github.com/gin-gonic/gin"
	"ppIm/ws"
)

func Test(ctx *gin.Context) {
	ws.SendToUser(100, "test", 200, "hello", gin.H{})
}
