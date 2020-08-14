package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ppIm/router"
)

func StartServer() {
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	r := gin.Default()
	router.SetRouter(r)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
