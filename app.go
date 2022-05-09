package main

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/routes"
	"go-webapp/utils"
	"log"
	"net/http"
)

func main() {
	models.TestConnection()

	//versão para desenvolvimento
	port := "8000"

	// versão para produção
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	fmt.Println("Not port specified")
	// 	os.Exit(1)
	// }

	fmt.Printf("Listening Port %s", port)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
