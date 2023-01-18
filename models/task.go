package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string
	TimeSpent string
	DueDate   string
	Current   bool
	Completed bool
	UserID    uint
}
