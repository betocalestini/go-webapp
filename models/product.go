package models

import "fmt"

type Product struct {
	Id       uint64   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Quantity int      `json:"quantity"`
	Amount   float64  `json:"amount"`
	Category Category `json:"category"`
}

func GetProducts() ([]Product, error) {
	con := Connect()
	defer con.Close()
	sql := "select c.id, c.description, p.id, p.nome, p.Price, p.Quantity, p.Amount from products as p inner join category as c on c.id = p.category"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var products []Product
	for rs.Next() {
		var product Product
		err := rs.Scan(&product.Category.Id,
			&product.Category.Description,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func SearchProducts(search string) ([]Product, error) {
	search = fmt.Sprintf("%%%s%%", search)
	con := Connect()
	defer con.Close()
	sql := `select c.id, LOWER(c.description), p.id, LOWER(p.nome), p.Price, p.Quantity, p.Amount 
										from products as p 
										inner join category as c on c.id = p.category
										where LOWER(c.description) like LOWER($1) or LOWER(p.nome) like LOWER($2) `
	stmt, err := con.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rs, err := stmt.Query(search, search)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var products []Product
	for rs.Next() {
		var product Product
		err := rs.Scan(&product.Category.Id,
			&product.Category.Description,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
