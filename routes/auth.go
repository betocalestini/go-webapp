package routes

import (
	"go-webapp/auth"
	"go-webapp/utils"
	"log"
	"net/http"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	user, err := auth.Signin(email, password)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	utils.ToJson(w, user)
}
