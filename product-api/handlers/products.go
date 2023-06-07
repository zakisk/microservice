package handlers

import (
	"fmt"
	"log"

	"github.com/zakisk/microservice/product-api/data"
)

// struct that implements Handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

type KeyProduct struct{}

// NewProducts creates a new Products handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
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
