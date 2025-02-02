package main

import (
	"fmt"
	"testing"

	"github.com/estnafinema0/go-microservices/product-api/sdk/client"
	"github.com/estnafinema0/go-microservices/product-api/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	//t.Fail()
}

func TestOurClientSimple(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	c.Products.ListProducts(params)
}
