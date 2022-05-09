package routes

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/utils"
	"net/http"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Status: Internal Server Error"))
	}
	products, err := models.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Status: Internal Server Error"))
	}

	utils.ExecuteTemplate(w, "home.html", struct {
		Categories []models.Category
		Products   []models.Product
	}{
		Categories: categories,
		Products:   products,
	})
}

func HomePostHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(products)
	count := len(products)
	var html string = ""
	if count > 0 {
		html += "<table class='table table-hover'>"
		html += `<thead>
								<tr>
									<th scope="col">Id</th>
									<th scope="col">Categoria</th>
									<th scope="col">Nome</th>
									<th scope="col">Preço</th>
									<th scope="col">Quantidade</th>
									<th scope="col">Total</th>
								</tr>
						</thead>
						<tbody>`
		for _, p := range products {
			html += fmt.Sprintf(`
												<tr>
                           <td>%d</td>
                           <td>%s</td>
                           <td>%s</td>
                           <td>%.2f</td>
                           <td>%d</td>
                           <td>%.2f</td>
                        </tr>
												`, p.Id, p.Category.Description, p.Name, p.Price, p.Quantity, p.Amount)

		}
		html += "</tbody>"
		html += "</table>"
	} else {
		html += fmt.Sprintf(`<p class='alert alert-info'>Nada encontrado com <code><strong>"%s"</strong></code></p>`, search)
	}
	w.Write([]byte(html))
}
