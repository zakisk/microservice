package handlers

import (
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// Responses:
//  200: productResponse
//  500: internalServerError

func (p *Products) ListProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("Products endpoint is called")
	rw.Header().Add("Content-Type", "application/json")

	currency := r.URL.Query().Get("currency")

	lp, err := p.productsDB.GetProducts(currency)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

	err = data.ToJSON(lp, rw)
	if err != nil {
		p.l.Error("Serializing products", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

}

// swagger:route GET /products/{id} products listSingleProducts
// Returns a list of products
// Responses:
//  200: productResponse
//	404: notFound
//  500: internalServerError

func (p *Products) ListSingelProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getProductId(r)
	currency := r.URL.Query().Get("currency")

	prod, err := p.productsDB.GetProductById(id, currency)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.l.Error("fetching product", "error", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Error("[ERROR] fetching product", "error", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
}
