package models

import "gorm.io/gorm"

type Tab struct {
	gorm.Model
	Title  string
	UserID uint
}
