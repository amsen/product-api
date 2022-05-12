package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello handler called...")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Panic("Error reading the body!")
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Hello %s\n", d)

}
