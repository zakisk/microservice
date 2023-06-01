package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// model class for product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var ErrorProductNotFound = fmt.Errorf("Product not found of given id")

// get product struct from JSON
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Products []*Product

// return static products
func GetProducts() Products {
	return productList
}

func AddProduct(prod *Product) {
	prod.ID = getNextID()
	productList = append(productList, prod)
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	productList[pos] = prod
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for pos, p := range productList {
		if p.ID == id {
			return p, pos, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

func getNextID() int {
	last := productList[len(productList)-1]
	return last.ID + 1
}

// encodes products and writes them directly to http.ResponseWriter
func (p *Products) ToJSON(wr http.ResponseWriter) error {
	e := json.NewEncoder(wr)
	return e.Encode(p)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Lifeboy",
		Description: "a soap for bath",
		Price:       15,
		SKU:         "abc123",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Samsung Galaxy S10",
		Description: "Samsung's android mobile device",
		Price:       15000,
		SKU:         "def123",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}
