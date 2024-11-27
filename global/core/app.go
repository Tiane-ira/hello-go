package core

import (
	"context"
	"fmt"
	"hello-go/configs"
	"hello-go/zlog"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func StartServer(lc fx.Lifecycle, mux *http.ServeMux) {
	addr := fmt.Sprintf(":%d", configs.Get().App.Port)
	server := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := server.ListenAndServe()
				if err != nil {
					zlog.Logger.Error("server start failed", zap.Error(err))
					os.Exit(1)
				}
			}()
			zlog.Logger.Info("server start success", zap.String("addr", addr))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 优雅关闭
			return server.Shutdown(ctx)
		},
	})
}
