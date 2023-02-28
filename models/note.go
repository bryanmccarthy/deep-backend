package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content" gorm:"type:text"`
	TaskID  uint   `json:"task_id"`
	UserID  uint   `json:"user_id"`
}
