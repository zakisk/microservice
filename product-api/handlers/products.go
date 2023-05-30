package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/zakisk/microservice/product-api/data"
)

// struct that gonna implement Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new Products handler with the given logger.
// It follows the dependency injection model to allow flexibility
// and increase testability by injecting the logger dependency.
// The logger can be replaced with a mock logger during testing.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(wr, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(wr, r)
		return
	}

	if r.Method == http.MethodPut {
		rx := regexp.MustCompile(`/([-0-9]+)`)
		g := rx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(wr, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(wr, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(wr, "Invalid URL", http.StatusBadRequest)
			return
		}
		p.l.Println("id is ", id)
		p.updateProduct(id, wr, r)

	}

	//catch methods
	wr.WriteHeader(http.StatusMethodNotAllowed)
}

// handles http GET request for products
func (p *Products) getProducts(wr http.ResponseWriter, r *http.Request) {
	p.l.Println("Products endpoint is called")
	lp := data.GetProducts()
	err := lp.ToJSON(wr)
	if err != nil {
		http.Error(wr, "Oops something went wrong", http.StatusInternalServerError)
	}
}

// handles http POST request method
func (p *Products) addProduct(wr http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(wr, "Unable to unmarshal data", http.StatusInternalServerError)

	}
	p.l.Printf("Product : %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, wr http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(wr, "Unable to unmarshal data", http.StatusInternalServerError)

	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(wr, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(wr, "Invalid id", http.StatusNotFound)
		return
	}
}
