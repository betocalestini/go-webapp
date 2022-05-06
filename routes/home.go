package routes

import (
	"go-webapp/utils"
	"net/http"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
