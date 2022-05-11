package routes

import (
	"go-webapp/models"
	"go-webapp/utils"
	"net/http"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	// utils.ToJson(w, user)
	_, err := models.NewUser(user)
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	http.Redirect(w, r, "/register", http.StatusFound)
}
