package routes

import (
	"go-webapp/utils"
	"net/http"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}
