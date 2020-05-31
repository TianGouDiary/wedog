package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// DB 数据库链接单例
var DB *xorm.Engine

// Database 在中间件中初始化mysql链接
func InitDatabase(connString string) {
	// username:password@tcp(ip:port)/dbname?charset=utf8
	fmt.Println("connString: ", connString)
	db, err := xorm.NewEngine("mysql", connString)
	if err != nil {
		panic(err)
		return
	}
	DB = db
	// 最大连接池为 20 个
	DB.SetMaxOpenConns(20)
	DB.ShowSQL(true)
	migration()
	return
}
