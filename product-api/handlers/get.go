package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	protos "github.com/zakisk/microservice/currency/protos/currency"
	"github.com/zakisk/microservice/product-api/data"
)

// swagger:route GET /products products listSingleProducts
// Returns a list of products
// Responses:
//  200: productResponse
//  500: internalServerError

func (p *Products) ListProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	rw.Header().Add("Content-Type", "application/json")
	lp := data.GetProducts()
	err := data.ToJSON(lp, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

}

// swagger:route GET /products/{id} products listProducts
// Returns a list of products
// Responses:
//  200: productsResponse
//	404: notFound
//  500: internalServerError

func (p *Products) ListSingelProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	rw.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid request", http.StatusBadRequest)
		return
	}

	prod, err := data.GetProductById(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["INR"]),
		Destination: protos.Currencies(protos.Currencies_value["USD"]),
	}

	res, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Fatal("[ERROR] while get currency rate", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	prod.Price = prod.Price * res.Rate

	err = data.ToJSON(prod, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
}
