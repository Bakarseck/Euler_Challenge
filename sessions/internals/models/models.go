package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}
