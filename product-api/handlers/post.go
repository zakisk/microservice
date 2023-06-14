package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// handles http POST request method
func (p *Products) POST(wr http.ResponseWriter, r *http.Request) {
	p.l.Debug("Handle GET Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Info("Product : %#v", prod)
	p.productsDB.AddProduct(prod)
}
