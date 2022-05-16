package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/amsen/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall the products list received from database!", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("Received Product: %#v\n", prod)
	// lp := data.GetProducts()
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Could not fetch the id to update, invalid id.", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

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

type KeyProduct struct{}

func (p *Products) ValidateProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		prod := data.NewProduct()
		// prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Could not decode the payload..", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
