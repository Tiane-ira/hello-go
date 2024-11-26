package core

import (
	"fmt"
	"hello-go/configs"
	"hello-go/router"
	"hello-go/zlog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func AppStart() {
	engine := router.InitRouters()

	addr := fmt.Sprintf(":%d", configs.Get().App.Port)
	s := initServer(addr, engine)

	zlog.Logger.Info("server run success on ", zap.String("addr", addr))
	err := s.ListenAndServe()
	if err != nil {
		zlog.Logger.Error(err.Error())
	}
}

func initServer(addr string, engine *gin.Engine) server {
	return &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
