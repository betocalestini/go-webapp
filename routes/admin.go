package routes

import (
	"go-webapp/models"
	"go-webapp/utils"
	"net/http"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	total := int64(len(products))
	utils.ExecuteTemplate(w, "admin.html", struct {
		Total    int64 `json:"total"`
		Products []models.Product
	}{
		Total:    total,
		Products: products,
	})
}
