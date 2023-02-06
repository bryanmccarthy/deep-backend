package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string
	TimeSpent uint
	Current   bool
	Completed bool
	UserID    uint
}
