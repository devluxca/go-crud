package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID 		uint `json:"id" gorm:"primary_key"`
	Name 	string `json:"name"`
	Age 	int `json:"age"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}