package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zakisk/microservice/product-api/data"
)

//swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from database of given id
// responses:
//  	201: noContentResponse
//		404: notFound
//		500: internalServerError
func (p *Products) DELETE(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "invalid request", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
}