package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// handles http Update request method
// swagger:route PUT /products products updateProduct
// Responses:
//	201: noContentResponse
func (p *Products) Update(wr http.ResponseWriter, r *http.Request) {

	prod := (r.Context().Value(KeyProduct{}).(data.Product))
	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		http.Error(wr, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(wr, "Invalid id", http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusNoContent)
}
