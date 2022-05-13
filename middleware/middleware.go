package middleware

import (
	"go-webapp/sessions"
	"net/http"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		_, ok := session.Values["USERID"]
		if !ok || session.Values["USERID"] == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
