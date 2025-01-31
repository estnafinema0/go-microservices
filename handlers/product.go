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

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch the products from the database
	lp := data.GetProducts()
	e := lp.ToJSON(rw)

	if e != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	e := prod.FromJSON(r.Body)

	if e != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Product) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	prod := &data.Product{}
	e := prod.FromJSON(r.Body)

	if e != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}

	e = data.UpdateProduct(id, prod)
	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
