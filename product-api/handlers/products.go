// Package classification awesome.
//
// Documentation of our awesome API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:9090
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package handlers


import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

// struct that implements Handler
type Products struct {
	l *log.Logger
}


// A list of products being returned in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the database
	// in: body
	Body []data.Product
}


// swagger:response noContent
type productNoContent struct {}

// swagger:parameters deleteProduct
type productIDParameter struct {
	// The id of the product to delete it from the database
	// in: path
	// required: true
	ID int `json: "id"`
}

type KeyProduct struct{}

// NewProducts creates a new Products handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}





// Middleware to validate Product
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product")
			http.Error(wr, "Unable to unmarshal data", http.StatusBadRequest)
			return
		}

		//validate the product we've got in request
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(
				wr,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		//add prod to this context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		//passing control to the next middleware or if there no then to final handler
		next.ServeHTTP(wr, r)
	})
}
