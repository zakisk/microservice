package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse

func (p *Products) GET(wr http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	lp := data.GetProducts()
	err := lp.ToJSON(wr)
	if err != nil {
		http.Error(wr, "Oops something went wrong", http.StatusInternalServerError)
	}
}
