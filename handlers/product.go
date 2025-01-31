package handlers

import (
	"log"
	"net/http"

	"github.com/estnafinema0/go-microservices/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	e := lp.ToJSON(rw)

	if e != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
