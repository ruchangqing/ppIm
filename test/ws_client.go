package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"ppIm/app/websocket"
)

func main() {
	// 并发测试
	c := make(chan int, 1)
	for i := 1; i <= 10000; i++ {
		go client()
	}
	select {
	case <-c:
		return
	}
}

func client() {
	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		log.Fatal("连接错误：", err)
	}
	defer c.Close()
	fmt.Println("已连接")

	c.WriteJSON(websocket.Message{Cmd: 3, Body: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdCI6MTYxOTE0Mzg0OSwiaWQiOjV9.KQ7dOv6bE_fP5NpMehziesFMsZXDAdVrbYBHyZROw40"})

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case <-done:
			return
		}
	}
}
