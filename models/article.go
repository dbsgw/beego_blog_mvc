package models

import (
	"gorm.io/gorm"
)

// Article 的定义
type Article struct {
	gorm.Model
	Title       string `gorm:"comment:标题"`                   // 标题
	Content     string `gorm:"type:longtext;comment:姓名"`     // mk内容
	ContentHtml string `gorm:"type:longtext;comment:html内容"` // html内容
	SortId      int64  // 分类Id
}

// TableName 重命名表
func (u *Article) TableName() string {
	return "Article"
}
