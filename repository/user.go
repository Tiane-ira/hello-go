package repository

import (
	"hello-go/core/db/paginator"
	"hello-go/core/db/sorter"
	"hello-go/domain"
	"hello-go/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) ListUser(req *domain.ListUserReq) (userList []model.CsUser, err error) {
	tx := u.db.Model(&model.CsUser{})
	if req.Start != nil && req.End != nil {
		tx.Where("created_at >= ?", req.Start).Where("created_at <= ?", req.End)
	}
	err = tx.Find(&userList).Error
	return
}

func (u *UserRepository) PageListUser(req *domain.ListUserPageReq) (pageData paginator.Page[model.CsUser], err error) {
	tx := u.db.Model(&model.CsUser{})
	if !req.Start.IsZero() && !req.End.IsZero() {
		tx.Where("created_at >= ?", req.Start).Where("created_at <= ?", req.End)
	}
	// 排序
	sorter.Sort(tx, req.Sort)
	// 分页
	pageData = paginator.Page[model.CsUser]{CurrentPage: req.CurPage, PageSize: req.Size}
	pageData.WithNoPage(req.NoPage)
	err = pageData.SelectPages(tx)
	return
}

func (u *UserRepository) GetById(id uint) (user model.CsUser, err error) {
	err = u.db.Model(&model.CsUser{}).Where("id = ?", id).First(&user).Error
	return
}

func (u *UserRepository) DeleteById(id uint) (err error) {
	err = u.db.Delete(&model.CsUser{}, id).Error
	return
}

func (u *UserRepository) SaveOrUpdate(user *model.CsUser) (err error) {
	if user.ID == 0 {
		err = u.db.Model(&model.CsUser{}).Create(user).Error
	} else {
		err = u.db.Model(&model.CsUser{}).Where("id = ?", user.ID).Updates(user).Error // 默认忽略零值更新
		// err = u.db.Model(&model.CsUser{}).Select("name", "age").Where("id = ?", user.ID).Updates(user).Error
	}
	return
}
