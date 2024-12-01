package app

import (
	"fmt"
	"hello-go/configs"
	"hello-go/zlog"
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func NewGin() (*gin.Engine, *gin.RouterGroup, *http.ServeMux) {

	pprofPath := "/debug/pprof"
	mux := http.NewServeMux()
	{
		mux.HandleFunc(fmt.Sprintf("%s/", pprofPath), pprof.Index)
		mux.HandleFunc(fmt.Sprintf("%s/cmdline", pprofPath), pprof.Cmdline)
		mux.HandleFunc(fmt.Sprintf("%s/profile", pprofPath), pprof.Profile)
		mux.HandleFunc(fmt.Sprintf("%s/trace", pprofPath), pprof.Trace)
		mux.HandleFunc(fmt.Sprintf("%s/symbol", pprofPath), pprof.Symbol)
	}

	e := gin.New()
	// 处理全局中间件
	e.Use(gin.Recovery())
	apiPrefix := configs.Get().App.ApiPrefix
	zlog.Logger.Info(fmt.Sprintf("get api prefix: %s", apiPrefix))
	rg := e.Group(apiPrefix)
	{
		rg.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	mux.Handle("/", e)

	return e, rg, mux
}
