package main

import (
	"ppIm/framework"
	_ "ppIm/global"
)

func main() {
	framework.LoadConfig()
	framework.LoadService()
	framework.StartServer()
}