package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeGetHandler).Methods("GET")
	return r
}
