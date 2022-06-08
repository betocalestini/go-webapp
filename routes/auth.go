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

	_, isAuth := sessions.IsLogged(r)
	if isAuth {
		http.Redirect(w, r, "/admin", 302)
	}

	session, _ := sessions.Store.Get(r, "session")
	var message string
	var alert string
	var active string

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
	delete(session.Values, "MESSAGE")
	delete(session.Values, "ALERT")
	delete(session.Values, "ACTIVE")
	delete(session.Values, "USERID")

	session.Save(r, w)

	utils.ExecuteTemplate(w, "login.html", struct {
		Alert utils.Alert
	}{Alert: utils.NewAlert(message, alert, active)})
}
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	user, err := auth.Signin(email, password)
	checkErrAuthenticate(err, w, r, user)
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request, user models.User) {
	session, _ := sessions.Store.Get(r, "session")

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

			http.Redirect(w, r, "/", http.StatusFound)
			return
		default:
			utils.InternalServerError(w)
			return
		}
	}
	session.Values["USERID"] = user.Id
	session.Save(r, w)
	http.Redirect(w, r, "/home", http.StatusFound)
	// w.Write([]byte("Usuário logado com sucesso!"))
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "MESSAGE")
	delete(session.Values, "ALERT")
	delete(session.Values, "ACTIVE")
	delete(session.Values, "USERID")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
