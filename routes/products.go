package routes

import (
	"fmt"
	"go-webapp/models"
	"go-webapp/sessions"
	"go-webapp/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func productsGetHandler(w http.ResponseWriter, r *http.Request) {
	allProducts, allUsers, err := models.LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	message, alert, active := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "products.html", struct {
		AllUsers    int64 `json:"AllUsers"`
		AllProducts int64 `json:"AllProducts"`
		Alert       utils.Alert
	}{
		AllUsers:    allUsers,
		AllProducts: allProducts,
		Alert:       utils.NewAlert(message, alert, active),
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
									<th scope="col" class="text-center">Editar</th>
									<th scope="col" class="text-center">Deletar</th>
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
                           <td class="text-center"><a href="/product-edit?productId=%d">
													 <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pen" viewBox="0 0 16 16">
													 <path d="m13.498.795.149-.149a1.207 1.207 0 1 1 1.707 1.708l-.149.148a1.5 1.5 0 0 1-.059 2.059L4.854 14.854a.5.5 0 0 1-.233.131l-4 1a.5.5 0 0 1-.606-.606l1-4a.5.5 0 0 1 .131-.232l9.642-9.642a.5.5 0 0 0-.642.056L6.854 4.854a.5.5 0 1 1-.708-.708L9.44.854A1.5 1.5 0 0 1 11.5.796a1.5 1.5 0 0 1 1.998-.001zm-.644.766a.5.5 0 0 0-.707 0L1.95 11.756l-.764 3.057 3.057-.764L14.44 3.854a.5.5 0 0 0 0-.708l-1.585-1.585z"/>
												 </svg></a></td>
                           <td class="text-center text-danger"><a href="/product-delete?productId=%d&confirm=true" class="text-danger"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
													 <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
													 <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
												 </svg>
													 </a></td>
                        </tr>
												`, p.Id, p.Category.Description, p.Name, p.Price, p.Quantity, p.Amount, p.Id, p.Id)

		}
		html += "</tbody>"
		html += "</table>"
	} else {
		// html += fmt.Sprintf(`<p class='alert alert-info'>Nada encontrado com <code><strong>"%s"</strong></code></p>`, search)
		html += fmt.Sprintf(`0_<p class='alert alert-info'>Nada encontrado com <code><strong>"%s"</strong></code></p>`, search)
	}
	w.Write([]byte(html))
}

func productCreateGetHandler(w http.ResponseWriter, r *http.Request) {
	allProducts, allUsers, err := models.LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	categories, err := models.GetCategories()
	if err != nil {
		utils.InternalServerError(w)
		return
	} else {
		utils.ExecuteTemplate(w, "product-create.html", struct {
			Categories  []models.Category
			AllUsers    int64 `json:"AllUsers"`
			AllProducts int64 `json:"AllProducts"`
		}{
			Categories:  categories,
			AllUsers:    allUsers,
			AllProducts: allProducts,
		})
	}
}

func productCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := verifyInputProduct(r)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	_, err = models.NewProduct(product)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	sessions.Message("Novo produto adicionado", "success", "", r, w)
	fmt.Println(sessions.Store.Get(r, "session"))
	http.Redirect(w, r, "/products", 302)
}

func verifyInputProduct(r *http.Request) (models.Product, error) {
	r.ParseForm()
	var err error = nil
	var product models.Product
	product.Id, err = strconv.ParseUint(r.PostForm.Get("id"), 10, 64)
	if err != nil {
		return models.Product{}, err
	}
	product.Name = r.PostForm.Get("name")
	price := strings.ReplaceAll(r.PostForm.Get("price"), ",", ".")
	product.Price, err = strconv.ParseFloat(price, 64)
	if err != nil {
		return models.Product{}, err
	}
	product.Quantity, err = strconv.Atoi(r.PostForm.Get("quantity"))
	if err != nil {
		return models.Product{}, err
	}
	product.Category.Id, err = strconv.Atoi(r.PostForm.Get("category"))
	if err != nil {
		return models.Product{}, err
	}
	quantity := float64(product.Quantity)
	product.Amount = product.Price * quantity
	return product, nil
}

func productEditGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	productId, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)
	product, err := models.GetProductById(productId)
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	priceFormat := fmt.Sprintf("%.2f", product.Price)
	priceFormat = strings.ReplaceAll(priceFormat, ".", ",")

	allProducts, allUsers, err := models.LoadData()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	categories, err := models.GetCategories()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	utils.ExecuteTemplate(w, "product-edit.html", struct {
		Categories  []models.Category
		Product     models.Product
		PriceFormat string
		AllUsers    int64 `json:"AllUsers"`
		AllProducts int64 `json:"AllProducts"`
	}{
		Categories:  categories,
		Product:     product,
		PriceFormat: priceFormat,
		AllUsers:    allUsers,
		AllProducts: allProducts,
	})
}

func productEditPostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := verifyInputProduct(r)
	fmt.Println(product)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	rows, err := models.UpdateProduct(product)
	if err != nil {
		log.Println(err)
		utils.InternalServerError(w)
		return
	}
	if rows > 0 {
		sessions.Message(fmt.Sprintf(`O Produto "%d" foi atualizado com sucesso`, product.Id), "success", "", r, w)
		fmt.Println(sessions.Store.Get(r, "session"))
		http.Redirect(w, r, "/products", 302)
	} else {
		sessions.Message(fmt.Sprintf("Não foi possível atualizar o produto #%d", product.Id), "danger", "", r, w)
		fmt.Println(sessions.Store.Get(r, "session"))
		http.Redirect(w, r, "/products", 302)
	}
}

func productDeleteGetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	productId, _ := strconv.ParseUint(keys.Get("productId"), 10, 64)
	ok, _ := strconv.ParseBool(keys.Get("confirm"))
	if !ok {
		sessions.Message(fmt.Sprintf("Não foi possível deletar o produto #%d", productId), "danger", "", r, w)
		fmt.Println(sessions.Store.Get(r, "session"))
		http.Redirect(w, r, "/products", 302)
	}
	rows, err := models.DeleteProduct(productId)
	if err != nil {
		log.Println(err)
		sessions.Message(fmt.Sprintf("Não foi possível deletar o produto #%d", productId), "danger", "", r, w)
		fmt.Println(sessions.Store.Get(r, "session"))
		http.Redirect(w, r, "/products", 302)
	}
	if rows > 0 {
		sessions.Message(fmt.Sprintf("O Produto %d foi deletado permanentemente", productId), "success", "", r, w)
		fmt.Println("oaisoidoasidjh")
		fmt.Println(sessions.Store.Get(r, "session"))
		http.Redirect(w, r, "/products", 302)
	}
}
