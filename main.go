package main

import (
	"fmt"
	"ppIm/boot"
)

func main() {
	// 全局错误处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//服务启动
	boot.LoadConfig()
	boot.LoadService()
	boot.StartServer()
}
