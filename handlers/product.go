package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/estnafinema0/go-microservices/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// Expect the id in the URI
		re := regexp.MustCompile(`/([0-9]+)`)
		matches := re.FindAllStringSubmatch(r.URL.Path, -1)
		if len(matches) != 1 {
			p.l.Println("Invalid URI: more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(matches[0]) != 2 {
			p.l.Println("Invalid URI: more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := matches[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI: unable to convert to number")
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}
		p.l.Println("Got id:", id)
		p.UpdateProduct(id, rw, r)
		return
	}

	// Catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
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
