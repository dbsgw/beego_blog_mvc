package models

import (
	"gorm.io/gorm"
)

// User 的定义
type User struct {
	gorm.Model
	Username string `gorm:"comment:用户名"` // 用户名
	Password string `gorm:"comment:密码"`  // 密码
}

// TableName 重命名表
func (u *User) TableName() string {
	return "User"
}
