package routes

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/sessions"
	"go-webapp/utils"
	"net/http"
)

// func registerGetHandler(w http.ResponseWriter, r *http.Request) {
// 	utils.ExecuteTemplate(w, "login.html", nil)
// }

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	// utils.ToJson(w, user)
	_, err := models.NewUser(user)
	checkErrRegister(err, w, r)

	http.Redirect(w, r, "/", http.StatusFound)
}

func checkErrRegister(err error, w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	message := "Cadastro realizado com sucesso!"
	if err != nil {
		switch err {
		case models.ErrRequiredFirstName,
			models.ErrRequiredLastName,
			models.ErrRequiredEmail,
			models.ErrRequiredPassword:
			message = fmt.Sprintf("%s", err)
			session.Values["ALERT"] = "danger"
		default:
			session.Values["ALERT"] = "danger"
			utils.InternalServerError(w)
			return
		}
		session.Values["MESSAGE"] = message
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
	session.Values["MESSAGE"] = message
	session.Values["ALERT"] = "success"
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
