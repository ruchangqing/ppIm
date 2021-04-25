package main

import (
	"fmt"
	"ppIm/framework"
	_ "ppIm/global"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	framework.LoadConfig()
	framework.LoadService()
	framework.StartServer()
}
