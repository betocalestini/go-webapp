package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeGetHandler).Methods("GET")
	r.HandleFunc("/home", homePostHandler).Methods("POST")

	r.HandleFunc("/", loginGetHandler).Methods("GET")
	r.HandleFunc("/", loginPostHandler).Methods("POST")

	// r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r
}
