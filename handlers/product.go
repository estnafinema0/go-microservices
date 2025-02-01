// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta

package handlers

import (
	"log"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productsResponces

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//swagger:parameters deleteProduct
type productIDParameter struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}
