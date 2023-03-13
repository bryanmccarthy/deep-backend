package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Difficulty uint   `json:"difficulty"`
	DueDate    string `json:"due_date"`
	Completed  bool   `json:"completed"`
	Progress   uint   `json:"progress"`
	TimeSpent  uint   `json:"time_spent"`
	UserID     uint   `json:"user_id"`
}
