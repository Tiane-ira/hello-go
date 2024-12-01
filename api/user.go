package api

import (
	"hello-go/core/app"
	"hello-go/domain"
	"hello-go/service"
	"hello-go/zlog"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(rg *gin.RouterGroup, userService *service.UserService) {
	handler := &UserHandler{userService: userService}
	userGroup := rg.Group("/user")
	{
		userGroup.GET(app.BindQuery("/list", handler.ListUser))
		userGroup.GET(app.BindQuery("/pageList", handler.PageListUser))
		userGroup.GET(app.BindUriAndQuery("/:id", handler.GetById))
		userGroup.POST(app.BindJson("/save", handler.Save))
		userGroup.POST(app.BindUriAndQuery("/delete/:id", handler.Remove))
	}
}

func (u *UserHandler) ListUser(a *app.AppGin, req *domain.ListUserReq) error {
	zlog.Logger.Info("入参:", zap.String("start", req.Start.String()))
	userList, err := u.userService.List(req)
	if err != nil {
		return err
	}
	return a.R(userList)
}

func (u *UserHandler) PageListUser(a *app.AppGin, req *domain.ListUserPageReq) error {

	pageData, err := u.userService.PageList(req)
	if err != nil {
		return err
	}
	return a.R(pageData)
}

func (u *UserHandler) GetById(a *app.AppGin, req *domain.IdReq) error {
	user, err := u.userService.GetById(req.Id)
	if err != nil {
		return err
	}
	return a.R(user)
}

func (u *UserHandler) Save(a *app.AppGin, req *domain.UserReq) error {
	newUser, err := u.userService.Save(req)
	if err != nil {
		return err
	}
	return a.R(newUser)
}

func (u *UserHandler) Remove(a *app.AppGin, req *domain.IdReq) error {
	err := u.userService.Remove(req.Id)
	if err != nil {
		return err
	}
	return a.R(nil)
}
