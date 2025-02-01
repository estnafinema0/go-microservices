package handlers

import (
	"net/http"

	"github.com/estnafinema0/go-microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 		200: productsResponces

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch the products from the database
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
