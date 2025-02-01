package handlers

import (
	"net/http"
	"strconv"

	"github.com/estnafinema0/go-microservices/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products products updateProduct
// Update a products details
// responses:
// 		201: noContentResponse
//  404: errorResponse
//  422: errorValidation

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
