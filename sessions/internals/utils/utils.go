package utils

import (
	"sessions/internals/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, firstName, lastName, email, password string) error {

	user := models.User{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	return db.Create(&user).Error
}
