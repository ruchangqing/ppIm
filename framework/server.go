package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ppIm/router"
)

func StartServer() {
	httpServer()
	/*	wg := sync.WaitGroup{}
		wg.Add(1)
		go httpServer(&wg)
		wg.Wait()
	*/
}

func httpServer(/*wg *sync.WaitGroup*/) {
	//gin.SetMode(gin.ReleaseMode)
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	server := gin.Default()
	router.SetRouter(server)
	panic(server.Run(fmt.Sprintf("%s:%s", host, port)))
	//wg.Done()
}
