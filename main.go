package main

import (
	"context"
	"hello-go/api"
	"hello-go/configs"
	"hello-go/global/core"
	"hello-go/global/core/app"
	"hello-go/global/db"
	"hello-go/global/redis"
	"hello-go/zlog"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	zlog.InitJsonZap(
		zlog.WithFileRotate(configs.Get().App.LogFile),
	)

	app := fx.New(
		fx.Provide(app.NewGin, redis.NewClinet, db.NewMysqlDb),
		fx.Invoke(api.NewUserHandler),
		fx.Invoke(core.StartServer),
	)

	if err := app.Start(context.Background()); err != nil {
		zlog.Logger.Error("app start failed", zap.Error(err))
		return
	}
	zlog.Logger.Info("app start success")
	//	服务保持
	<-app.Done()
	zlog.Logger.Info("app exit")
}
