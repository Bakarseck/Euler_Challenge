package models

import (
	"net/http"

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

type EndPoint struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}
