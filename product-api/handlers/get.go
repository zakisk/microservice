package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// Responses:
//  200: productsResponse
//  500: internalServerError

func (p *Products) GET(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	rw.Header().Add("Content-Type", "application/json")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

}