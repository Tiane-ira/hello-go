package model

import (
	"hello-go/core/app"

	"gorm.io/gorm"
)

// ObjBase 基础模型，每张表都有的字段
type ObjBase struct {
	ID        uint           `gorm:"primarykey" json:"id"`            // 主键ID
	CreatedAt *app.DateTime  `json:"createdAt" grom:"autoCreateTime"` // 创建时间
	UpdatedAt *app.DateTime  `json:"updatedAt"`                       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                  // 删除时间
}
