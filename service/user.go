package service

import (
	"fmt"
	"hello-go/core/db/paginator"
	"hello-go/domain"
	"hello-go/model"
	"hello-go/repository"
	"hello-go/utils/redis"
	"hello-go/zlog"

	"go.uber.org/zap"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) List(req *domain.ListUserReq) (userList []model.CsUser, err error) {
	return u.userRepo.ListUser(req)
}

func (u *UserService) PageList(req *domain.ListUserPageReq) (pageData paginator.Page[model.CsUser], err error) {
	return u.userRepo.PageListUser(req)
}

func (u *UserService) GetById(id uint) (user model.CsUser, err error) {
	err = redis.GetObj(fmt.Sprintf("user::%d", id), &user)
	if err != nil {
		zlog.Logger.Error("get user from redis failed", zap.Error(err))
	}
	if user.ID != 0 {
		zlog.Logger.Info("get user from redis success")
		return
	}
	user, err = u.userRepo.GetById(id)
	if err != nil {
		return
	}
	err = redis.Set(fmt.Sprintf("user::%d", id), &user)
	return
}

func (u *UserService) Remove(id uint) (err error) {
	return u.userRepo.DeleteById(id)
}

func (u *UserService) Save(req *domain.UserReq) (newUser model.CsUser, err error) {
	user := &model.CsUser{ObjBase: model.ObjBase{ID: req.Id}, Name: req.Name, Age: req.Age}
	err = u.userRepo.SaveOrUpdate(user)
	if err != nil {
		return
	}
	newUser, err = u.GetById(user.ID)
	return
}
