package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}


// NewGoodbye creates a new Goodbye handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}


func (g *Goodbye) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(wr, "Oops! something went wrong", http.StatusBadRequest)
	}
	
	g.l.Printf("Goodbye %s", data)
	fmt.Fprintf(wr, "Goodbye %s", data)
}