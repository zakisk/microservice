package data

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	protos "github.com/zakisk/microservice/currency/protos/currency"

	"github.com/go-playground/validator"
)

// model class for product
// swagger:model
type Product struct {
	// id of product
	//
	// required: true
	ID int `json:"id"`
	// name of product
	//
	// required: true
	// min length: 2
	Name string `json:"name" validate:"required"`
	// description of product
	//
	Description string `json:"description"`
	// price of product
	//
	// required: true
	// min: 1
	Price float64 `json:"price" validate:"gt=0"`
	// unique identifier of the product
	//
	//required: true
	//pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
	//timestamp when the product is created
	//
	CreatedOn string `json:"-"`
	// timestamp when the product is updated
	//
	UpdatedOn string `json:"-"`
	// timestamp when the product is deleted
	//
	DeletedOn string `json:"-"`
}

// ErrProductNotFound is raised when product is not found in database
var ErrProductNotFound = fmt.Errorf("Product not found")

type Products []*Product


// type to handle product CURD operations
type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}


func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{c, l}
}


// return static products
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil	
	}

	rate, err := p.getRate(currency)
	p.log.Error("rate", rate)
	if err != nil {
		p.log.Error("Unable to get rate", currency, "currency", "error", err)
		return nil, err
	}

	prods := Products{}
	for _, p := range productList {
		dp := *p
		dp.Price = dp.Price * rate
		prods = append(prods, &dp)
	}

	return prods, nil
}

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}
	return validate.Struct(p)
}

func (p *ProductsDB) AddProduct(prod Product) {
	maxID := productList[len(productList) - 1].ID
	prod.ID = maxID + 1
	productList = append(productList, &prod)
}

// Updates the product in the database
func (p *ProductsDB) UpdateProduct(prod Product) error {
	pos := findIndexByProductID(prod.ID)
	if pos == -1 {
		return ErrProductNotFound
	}

	productList[pos] = &prod

	return nil
}

func (p *ProductsDB) GetProductById(id int, currency string) (Product, error) {
	i := findIndexByProductID(id)
	if 1 == -1 {
		return Product{}, ErrProductNotFound
	}

	if currency == "" {
		return *productList[i], nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", currency, "currency", "error", err)
		return Product{}, err
	}

	dp := *productList[i]
	dp.Price = dp.Price * rate
	return dp, nil
}

func (p *ProductsDB) DeleteProduct(id int) error {
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


func (p *ProductsDB) getRate(destination string) (float64, error) {
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["USD"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	res, err := p.currency.GetRate(context.Background(), rr)
	return res.Rate, err
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
