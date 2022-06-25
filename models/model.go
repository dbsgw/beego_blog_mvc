package models

import "time"

type Model struct {
	Id        uint      `gorm:"primarykey;comment:自增主键"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`
	DeletedAt time.Time `gorm:"index;comment:删除时间"`
}
