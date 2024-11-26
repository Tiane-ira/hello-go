package main

import (
	"hello-go/configs"
	"hello-go/global/core"
	"hello-go/global/db"
	"hello-go/global/redis"
	"hello-go/zlog"
)

func main() {

	// 初始化日志
	zlog.InitJsonZap(
		zlog.WithFileRotate(configs.Get().App.LogFile),
	)
	// 初始化redis
	redis.InitClinet()
	// 初始化mysql
	db.InitMysql()
	if db.MysqlClient != nil {
		db.RegisterTables()
		db, _ := db.MysqlClient.DB()
		defer db.Close()
	}

	// 启动服务
	core.AppStart()
}
