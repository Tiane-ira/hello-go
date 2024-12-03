package main

import (
	"context"
	"hello-go/api"
	"hello-go/configs"
	"hello-go/core/app"
	"hello-go/core/db"
	"hello-go/repository"
	"hello-go/service"
	"hello-go/utils/redis"
	"hello-go/zlog"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	zlog.InitJsonZap(
		zlog.WithFileRotate(configs.Get().App.LogFile),
	)

	a := fx.New(
		fx.Invoke(db.RegisterTables, redis.InitRedis),
		fx.Provide(app.NewGin, db.NewMysqlDb),
		fx.Provide(service.NewUserService),
		fx.Provide(repository.NewUserRepository),
		fx.Invoke(api.NewUserHandler),
		fx.Invoke(app.StartServer),
	)

	if err := a.Start(context.Background()); err != nil {
		zlog.Logger.Error("app start failed", zap.Error(err))
		return
	}
	zlog.Logger.Info("app start success")
	//	服务保持
	<-a.Done()
	zlog.Logger.Info("app exit")
}
