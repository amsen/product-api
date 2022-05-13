package data

import (
	"encoding/json"
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
