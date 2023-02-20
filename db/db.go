package db

import (
	"github.com/bryanmccarthy/deep-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Note{})

	return db
}
