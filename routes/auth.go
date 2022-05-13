package routes

import (
	"fmt"
	"go-webapp/auth"
	"go-webapp/models"
	"go-webapp/sessions"
	"go-webapp/utils"
	"net/http"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	var message string
	var alert string
	var active string
	fmt.Println(session)
	if session.Values["MESSAGE"] != nil {
		message = fmt.Sprintf("%s", session.Values["MESSAGE"])
	} else {
		message = ""
	}
	if session.Values["ALERT"] != nil {
		alert = fmt.Sprintf("%s", session.Values["ALERT"])
	} else {
		alert = ""
	}
	if session.Values["ACTIVE"] != nil {
		active = fmt.Sprintf("%s", session.Values["ACTIVE"])
	} else {
		active = ""
	}

	session.Values["MESSAGE"] = ""
	session.Values["ALERT"] = ""
	session.Values["ACTIVE"] = ""
	session.Save(r, w)

	utils.ExecuteTemplate(w, "login.html", struct {
		Alert utils.Alert
	}{Alert: utils.NewAlert(message, alert, active)})
}
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	_, err := auth.Signin(email, password)
	if err != nil {
		checkErrAuthenticate(err, w, r)
	}
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	fmt.Println(session)
	if err != nil {
		fmt.Println(err)
		switch err {
		case auth.ErrEmptyFields,
			auth.ErrEmailNotFound,
			models.ErrInvalidEmail,
			auth.ErrInvalidPassword:
			session.Values["ALERT"] = "danger"
			session.Values["ACTIVE"] = ""
			session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
			session.Save(r, w)
			fmt.Println(session)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		default:
			utils.InternalServerError(w)
			return
		}
	}
	http.Redirect(w, r, "/home", http.StatusFound)
	// w.Write([]byte("Usuário logado com sucesso!"))
}
