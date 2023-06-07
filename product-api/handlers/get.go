package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse
//  500: internalServerError

func (p *Products) GET(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Oops something went wrong", http.StatusInternalServerError)
	}
}
