package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// handles http POST request method
func (p *Products) POST(wr http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Product : %#v", prod)
	data.AddProduct(&prod)
}
