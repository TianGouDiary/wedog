package main

import (
	"wedog/config"
	"wedog/server"
)

func main() {
	//加载配置
	config.Init()

	//装载路由
	r := server.NewRouter()
	r.Run(":80")
}
