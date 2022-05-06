package main

import (
	"fmt"
	"go-webapp/routes"
	"go-webapp/utils"
	"go-webapp/views/models"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	models.TestConnection()
	fmt.Println(models.GetCategories())
	fmt.Printf("Listening Port %s", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
