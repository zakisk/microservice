package main

import (
	"fmt"
	"testing"

	"github.com/zakisk/microservice/product-api/sdk/client"
	"github.com/zakisk/microservice/product-api/sdk/client/products"
)

func TestHttpClient(t *testing.T) {
	c := client.Default
	prods, err := c.Products.ListProducts(products.NewListProductsParams())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prods)
}