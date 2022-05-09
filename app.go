package main

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/routes"
	"go-webapp/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	models.TestConnection()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("Not port specified")
		os.Exit(1)
	}
	fmt.Printf("Listening Port %s", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
