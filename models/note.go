package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	TaskID  uint
	UserID  uint
}
