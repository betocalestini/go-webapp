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

	//comentar para produção
	// port := "8000"

	// comentar para desenvolvimento
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Not port specified")
		os.Exit(1)
	}

	fmt.Printf("Listening Port %s", port)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
