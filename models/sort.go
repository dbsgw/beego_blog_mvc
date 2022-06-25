package models

import (
	"gorm.io/gorm"
)

// Sort 的定义
type Sort struct {
	gorm.Model
	Title string `gorm:"comment:标题"` // 标题
}

// TableName 重命名表
func (u *Sort) TableName() string {
	return "Sort"
}
