package dao

import (
	"time"
)

//Model model default
type Model struct {
	ID       uint       `gorm:"column:id;primary_key;auto_increment;"`
	CreateAt *time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}

//ModelUser .
type ModelUser struct {
	Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Image    string `gorm:"column:image"`
}

//TableName .
func (a *ModelUser) TableName() string {
	return "user"
}
