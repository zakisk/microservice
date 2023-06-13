package data

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
)

// model class for product
// swagger:model
type Product struct {
	// id of product
	//
	// required: true
	ID          int     `json:"id"`
	// name of product
	//
	// required: true
	// min length: 2
	Name        string  `json:"name" validate:"required"`
	// description of product
	//
	Description string  `json:"description"`
	// price of product
	//
	// required: true
	// min: 1
	Price       float32 `json:"price" validate:"gt=0"`
	// unique identifier of the product
	//
	//required: true
	//pattern: [a-z]+-[a-z]+-[a-z]+
	SKU         string  `json:"sku" validate:"required,sku"`
	//timestamp when the product is created
	//
	CreatedOn   string  `json:"-"`
	// timestamp when the product is updated
	//
	UpdatedOn   string  `json:"-"`
	// timestamp when the product is deleted
	//
	DeletedOn   string  `json:"-"`
}

//ErrProductNotFound is raised when product is not found in database
var ErrProductNotFound = fmt.Errorf("Product not found")

type Products []*Product

// return static products
func GetProducts() Products {
	return productList
}

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}
	return validate.Struct(p)
}



func AddProduct(prod *Product) {
	prod.ID = getNextID()
	productList = append(productList, prod)
}

//Updates the product in the database
func UpdateProduct(prod Product) error {
	pos := findIndexByProductID(prod.ID)
	if pos == -1 {
		return ErrProductNotFound
	}

	productList[pos] = &prod
	return nil
}


func GetProductById(id int) (Product, error) {
	i := findIndexByProductID(id)
	if 1 == -1 {
		return Product{}, ErrProductNotFound
	}

	return *productList[i], nil
}


func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])
	return nil
}


func findIndexByProductID(id int) int {
	for pos, p := range productList {
		if p.ID == id {
			return pos
		}
	}

	return -1
}

func getNextID() int {
	last := productList[len(productList)-1]
	return last.ID + 1
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
