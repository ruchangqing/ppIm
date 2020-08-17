package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"ppIm/router"
	"sync"
)

func StartServer() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go httpServer(&wg)
	wg.Add(1)
	go socketIoServer(&wg)
	wg.Wait()
}

func httpServer(wg *sync.WaitGroup) {
	host := viper.GetString("http.host")
	port := viper.GetString("http.port")
	server := gin.Default()
	router.SetRouter(server)
	panic(server.Run(fmt.Sprintf("%s:%s", host, port)))
	wg.Done()
}

func socketIoServer(wg *sync.WaitGroup) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/", server)
	fmt.Println("[Socket.io] Server is running...")
	host := viper.GetString("socketio.host")
	port := viper.GetString("socketio.port")
	panic(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
	wg.Done()
}
