package routes

import (
	"fmt"
	"go-webapp/auth"
	"go-webapp/sessions"
	"go-webapp/utils"
	"net/http"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	// session, _ := sessions.Store.Get(r, "session")
	// var message string = ""
	// untypedMessage := session.Values["MESSAGE"]
	// message, _ = untypedMessage.(string)

	message, alert, active := sessions.Flash(r, w)

	utils.ExecuteTemplate(w, "login.html", struct {
		Alert utils.Alert
	}{Alert: utils.NewAlert(message, alert, active)})
}
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	_, err := auth.Signin(email, password)
	checkErrAuthenticate(err, w, r)
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	if err != nil {
		switch err {
		case auth.ErrInvalidEmail,
			auth.ErrInvalidPassword:
			session.Values["ALERT"] = "danger"
			session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
			session.Save(r, w)
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
