package service

import (
	"hello-go/core/db/paginator"
	"hello-go/domain"
	"hello-go/model"
	"hello-go/repository"
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
	return u.userRepo.GetById(id)
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
