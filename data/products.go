package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//defines the structure for an API product
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

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

func NewProduct() *Product {
	return &Product{}
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

// this is the dao here!
// this abstracts away how we can fetch the data from a database
// from the handlers
func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	i, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[i] = p
	return nil

}

var ErrproductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, prod := range productList {
		if id == prod.ID {
			return i, nil
		}
	}
	return 0, ErrproductNotFound
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          123,
		Name:        "Latte",
		Description: "Define latte here",
		Price:       3.14,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          456,
		Name:        "Espresso",
		Description: "Define latte here",
		Price:       6.28,
		SKU:         "def456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
