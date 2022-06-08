package routes

import (
	"go-webapp/models"
	"go-webapp/utils"
	"net/http"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	allProducts, allUsers, err := models.LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	utils.ExecuteTemplate(w, "admin.html", struct {
		AllUsers    int64 `json:"AllUsers"`
		AllProducts int64 `json:"AllProducts"`
	}{
		AllUsers:    allUsers,
		AllProducts: allProducts,
	})
}
