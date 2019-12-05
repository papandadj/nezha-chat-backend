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

//ModelUserImage .
type ModelUserImage struct {
	Model
	Name string `gorm:"column:name"`
	URL  string `gorm:"column:url"`
}

//TableName .
func (a *ModelUserImage) TableName() string {
	return "nezha_chat_common_user_image"
}
