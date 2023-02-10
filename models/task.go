package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title      string
	TimeSpent  uint
	Difficulty uint
	Completed  bool
	UserID     uint
}
