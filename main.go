package main

import (
	"fmt"
	"ppIm/framework"
	_ "ppIm/global"
)

func main() {
	// 全局错误处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//服务启动
	framework.LoadConfig()
	framework.LoadService()
	framework.StartServer()
}
