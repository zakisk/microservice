package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zakisk/microservice/product-api/data"
)

// struct that gonna implement Handler
type Products struct {
	l *log.Logger
}

type KeyProduct struct{}

// NewProducts creates a new Products handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// handles http GET request for products
func (p *Products) GetProducts(wr http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	lp := data.GetProducts()
	err := lp.ToJSON(wr)
	if err != nil {
		http.Error(wr, "Oops something went wrong", http.StatusInternalServerError)
	}
}

// handles http POST request method
func (p *Products) AddProduct(wr http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Product : %#v", prod)
	data.AddProduct(&prod) 
}

func (p *Products) UpdateProduct(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(wr, "Invalid id", http.StatusBadRequest)
	}

	prod := (r.Context().Value(KeyProduct{}).(data.Product))
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(wr, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(wr, "Invalid id", http.StatusBadRequest)
		return
	}
}

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product")
			http.Error(wr, "Unable to unmarshal data", http.StatusBadRequest)
			return
		}

		//add prod to this context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		//passing control to the next middleware or if there no then to final handler
		next.ServeHTTP(wr, r)
	})
}
