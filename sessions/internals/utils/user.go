package utils

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, username, firstName, lastName, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	statement, err := db.Prepare("INSERT INTO User (username, firstName, lastName, email, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(username, firstName, lastName, email, string(hashedPassword))
	return err
}

func CreateSession(db *sql.DB, username string) string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println("Error creating UUID: ", err.Error())
		return ""
	}

	statement, err := db.Prepare("INSERT INTO Sessions (token, username) VALUES (?, ?)")
	if err != nil {
		log.Println("Error preparing statement: ", err.Error())
		return ""
	}
	defer statement.Close()

	_, err = statement.Exec(uuid.String(), username)
	if err != nil {
		log.Println("Error executing statement: ", err.Error())
		return ""
	}

	return uuid.String()
}

func CheckUser(db *sql.DB, username, password string) bool {
	row := db.QueryRow("SELECT password FROM users WHERE username = ?", username)

	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
