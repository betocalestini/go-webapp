package routes

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/utils"
	"net/http"
)

func productsGetHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	total := int64(len(products))
	utils.ExecuteTemplate(w, "products.html", struct {
		Total    int64 `json:"total"`
		Products []models.Product
	}{
		Total:    total,
		Products: products,
	})
}

func productsPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencoded")
	r.ParseForm()
	//adequações para receber os dados do fetch
	teste := r.Form
	var search string
	for i := range teste {
		search = i
	}
	//adequações para receber os dados do fetch
	products, err := models.SearchProducts(search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Status: Internal Server Error"))
	}

	count := len(products)
	var html string = ""
	if count > 0 {
		html += "<table class='table table-hover' id='data-table'>"
		html += `<thead>
								<tr>
									<th scope="col" class="text-center">Id</th>
									<th scope="col" class="text-center">Categoria</th>
									<th scope="col" class="text-center">Nome</th>
									<th scope="col" class="text-center">Preço</th>
									<th scope="col" class="text-center">Quantidade</th>
									<th scope="col" class="text-center">Total</th>
								</tr>
						</thead>
						<tbody>`
		for _, p := range products {
			html += fmt.Sprintf(`
												<tr>
                           <td class="text-center">%d</td>
                           <td class="text-center">%s</td>
                           <td class="text-center">%s</td>
                           <td class="text-center">%.2f</td>
                           <td class="text-center">%d</td>
                           <td class="text-center">%.2f</td>
                        </tr>
												`, p.Id, p.Category.Description, p.Name, p.Price, p.Quantity, p.Amount)

		}
		html += "</tbody>"
		html += "</table>"
	} else {
		// html += fmt.Sprintf(`<p class='alert alert-info'>Nada encontrado com <code><strong>"%s"</strong></code></p>`, search)
		html += fmt.Sprintf(`0_<p class='alert alert-info'>Nada encontrado com <code><strong>"%s"</strong></code></p>`, search)
	}
	w.Write([]byte(html))
}
