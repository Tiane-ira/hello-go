package model

import (
	"time"

	"gorm.io/gorm"
)

// 基础模型，每张表都有的字段
type ObjBase struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
