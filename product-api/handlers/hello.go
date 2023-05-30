package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//Hello is a handler for endpoint "/"
type Hello struct {
	l *log.Logger
}

// NewHello creates a new Hello handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(wr, "Oopps! something went wrong", http.StatusBadRequest)
	}

	h.l.Printf("Hello %s", data)
	fmt.Fprintf(wr, "Hello %s", data)
}
