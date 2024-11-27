package api

import (
	"hello-go/global/core/app"
	"hello-go/zlog"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
}

type ListUserReq struct {
	Id int `form:"id"`
}

func NewUserHandler(rg *gin.RouterGroup) {
	handler := &UserHandler{}
	userGroup := rg.Group("/user")
	{
		userGroup.GET(app.BindQuery("/list", handler.ListUser))
		userGroup.GET(app.BindQuery("/list2", handler.ListUser2))
	}
}

func (u *UserHandler) ListUser(a *app.AppGin, req ListUserReq) error {
	zlog.Logger.Info("list user", zap.Any("req", req.Id))
	return a.R("ok")
}
func (u *UserHandler) ListUser2(a *app.AppGin, req *ListUserReq) error {
	zlog.Logger.Info("list user", zap.Any("req", req.Id))
	return a.R("ok")
}
