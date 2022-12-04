package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(&User{})

	if err != nil {
		return
	}

	DB = database
}