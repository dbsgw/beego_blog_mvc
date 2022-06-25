package models

// Sort 的定义
type Sort struct {
	Model
	Title string `gorm:"comment:标题"` // 标题
}

// TableName 重命名表
func (u *Sort) TableName() string {
	return "Sort"
}
