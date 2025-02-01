package handlers

import (
	"net/http"

	"github.com/estnafinema0/go-microservices/data"
)

// swagger:route POST /products products createProduct
// Create a new product
// responses:
// 		200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create a new product
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(prod)
}
