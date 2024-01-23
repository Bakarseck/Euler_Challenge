package handlers

import (
	"fmt"
	"net/http"
	"os"
	"sessions/internals/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.RenderPage("index", nil, w)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

	DB_PATH, SQL_FILE := os.Getenv("DB_PATH"), os.Getenv("SQL_FILE")
	db, err := utils.OpenDatabase(DB_PATH, SQL_FILE)

	if err != nil {
		fmt.Println("error database: ", err.Error())
	}

	r.ParseForm()
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	utils.CreateUser(db, username, firstname, lastname, email, password)
	utils.CreateSession(db, username)

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "",
	})
}
