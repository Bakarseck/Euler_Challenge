package handlers

import (
	"net/http"
	"sessions/internals/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.RenderPage("index", nil, w)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

	db := utils.ConnectDatabase()

	r.ParseForm()
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	utils.CreateUser(db, username, firstname, email, password)
}
