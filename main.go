package main

import (
	"context"
	"hello-go/api"
	"hello-go/configs"
	"hello-go/core/app"
	"hello-go/core/db"
	"hello-go/repository"
	"hello-go/service"
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
		fx.Provide(app.NewGin, db.NewMysqlDb),
		fx.Provide(service.NewUserService),
		fx.Provide(repository.NewUserRepository),
		fx.Invoke(api.NewUserHandler),
		fx.Invoke(app.StartServer),
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
