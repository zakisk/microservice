package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zakisk/microservice/product-api/data"
)

//swagger:route DELETE /products/{id} producs deleteProduct
// Deletes a product from database of given id
// responses:
//  	201: noContentResponse

func (p *Products) DELETE(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(wr, "invalid request", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(wr, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(wr, "Product not found", http.StatusInternalServerError)
		return
	}
}