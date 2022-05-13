package main

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/routes"
	"go-webapp/sessions"
	"go-webapp/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	models.TestConnection()

	//comentar para produção
	// port := "8000"

	// comentar para desenvolvimento
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Not port specified")
		os.Exit(1)
	}

	fmt.Printf("Listening Port %s \n", port)
	utils.LoadTemplates("views/*.html")
	sessions.SessionOptions("https://sheltered-citadel-21048.herokuapp.com", "/", 1800, true)
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
