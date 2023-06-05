package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/zakisk/microservice/product-api/data"
)

// handles http PUT request method
func (p *Products) PUT(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(wr, "Invalid id", http.StatusBadRequest)
	}

	prod := (r.Context().Value(KeyProduct{}).(data.Product))
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(wr, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(wr, "Invalid id", http.StatusBadRequest)
		return
	}
}
