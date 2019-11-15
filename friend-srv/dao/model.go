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

//ModelFriend .
type ModelFriend struct {
	Model
	UserID1 string `gorm:"column:user_id1"`
	UserID2 string `gorm:"column:user_id2"`
}

//TableName .
func (a *ModelFriend) TableName() string {
	return "friend"
}
