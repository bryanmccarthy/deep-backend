package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Difficulty uint   `json:"difficulty"`
	DueDate    string `json:"due_date"`
	Completed  bool   `json:"completed"`
	UserID     uint   `json:"user_id"`
}
