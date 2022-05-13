package routes

import (
	"go-webapp/utils"
	"net/http"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "admin.html", nil)
}
