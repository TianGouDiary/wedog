package config

import (
	"github.com/joho/godotenv"
	"os"
	_ "wedog/cache"
	"wedog/model"
)

func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 连接数据库
	model.InitDatabase(os.Getenv("MYSQL_DSN"))

	//TODO
	//cache.Redis()
}
