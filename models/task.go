package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string
	TimeSpent string
	Current   bool
	Completed bool
	UserID    uint
}
