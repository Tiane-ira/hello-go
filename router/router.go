package router

import (
	"fmt"
	"hello-go/api"
	"hello-go/configs"
	"hello-go/zlog"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	e := gin.Default()
	apiPrefix := configs.Get().App.ApiPrefix
	zlog.Logger.Info(fmt.Sprintf("get api prefix: %s", apiPrefix))
	publicGroup := e.Group(apiPrefix)
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.String(200, "ok")
		})
	}
	userGroup := e.Group(apiPrefix).Group("user")
	{
		userGroup.GET("/create", api.UserCreate)
	}
	return e
}
