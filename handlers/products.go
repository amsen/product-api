package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/amsen/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(w, r)
		return
	}

	if r.Method == http.MethodPut {
		uri := strings.Split(r.URL.String(), "/")
		id := uri[len(uri)-1]
		idval, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "failed to get id from url", http.StatusBadRequest)
		}
		p.l.Println(idval)
		p.updateProduct(idval, w, r)
		return
	}
	//catch all for all unimplemented methods
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall the products list received from database!", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
	prod := data.NewProduct()
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Could not decode the payload..", http.StatusBadRequest)
		return
	}
	p.l.Printf("Received Product: %#v\n", prod)
	// lp := data.GetProducts()
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	prod := data.NewProduct()
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Could not decode the payload..", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrproductNotFound {
		http.Error(w, data.ErrproductNotFound.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, data.ErrproductNotFound.Error(), http.StatusInternalServerError)
		return
	}
}
