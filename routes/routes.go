package routes

import (
	"go-webapp/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", loginGetHandler).Methods("GET")
	r.HandleFunc("/", loginPostHandler).Methods("POST")

	r.HandleFunc("/register", registerPostHandler).Methods("POST")

	r.HandleFunc("/home", middleware.AuthRequired(homeGetHandler)).Methods("GET")

	r.HandleFunc("/products", middleware.AuthRequired(productsGetHandler)).Methods("GET")
	r.HandleFunc("/products", middleware.AuthRequired(productsPostHandler)).Methods("POST")

	r.HandleFunc("/product-create", middleware.AuthRequired(productCreateGetHandler)).Methods("GET")
	r.HandleFunc("/product-create", middleware.AuthRequired(productCreatePostHandler)).Methods("POST")

	r.HandleFunc("/product-edit", middleware.AuthRequired(productEditGetHandler)).Methods("GET")
	r.HandleFunc("/product-edit", middleware.AuthRequired(productEditPostHandler)).Methods("POST")

	r.HandleFunc("/product-delete", middleware.AuthRequired(productDeleteGetHandler)).Methods("GET")

	r.HandleFunc("/admin", middleware.AuthRequired(adminGetHandler)).Methods("GET")
	r.HandleFunc("/logout", middleware.AuthRequired(logoutGetHandler)).Methods("GET")

	// r.HandleFunc("/register", registerGetHandler).Methods("GET")

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r
}
