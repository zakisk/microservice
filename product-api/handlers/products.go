package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/zakisk/microservice/product-api/data"
)

// struct that implements Handler
type Products struct {
	l          hclog.Logger
	v          *data.Validation
	productsDB *data.ProductsDB
}

type KeyProduct struct{}

// NewProducts creates a new Products handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewProducts(l hclog.Logger, v *data.Validation, pdb *data.ProductsDB) *Products {
	return &Products{l, v, pdb}
}

var ErrInvalidProductPath = fmt.Errorf("specified path is invalid, path should be /products/{id}")

type GenericError struct {
	// generic message for most of the errors
	Message string `json: "message"`
}

// ValidationError is a slice of json validation errors returned by validator
type ValidationError struct {
	Messages []string `json: "messages"`
}


// gets id from mux variabels
func getProductId(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	
	if err != nil {
		//this must not take place
		panic(err)
	}

	return id
}
  