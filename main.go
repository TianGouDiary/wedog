package main

import (
	"os"
	"wedog/config"
	"wedog/server"
)

func main() {
	//加载配置
	config.Init()

	//装载路由
	r := server.NewRouter()

	//默认监听80端口
	configPort := os.Getenv("SERVER_LISTON_PORT")
	if configPort == "" {
		configPort = "80"
	}
	err := r.Run(":" + configPort)
	if err != nil {
		panic(err)
	}
}
