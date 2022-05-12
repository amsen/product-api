package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amsen/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	l.Println("Starting up product api!")

	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	sig := <-sigchan
	l.Printf("Received terminate, graceful shutdown %s", sig)

	tctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tctx)

}
