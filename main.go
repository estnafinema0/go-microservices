package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/estnafinema0/go-microservices/handlers"
	"github.com/gorilla/mux"
)

// go run main.go
// curl -v -d "Hello World" localhost:9090/...
// curl -v -d "Nice" localhost:9090 | jq
// another method that migth not be allowed
// curl  localhost:9090 -XDELETE -v | jq
// POST with  json
// curl  localhost:9090 -d '{"id":1, "name": "tea", "description": "hehe"}' | jq
// PUT with  json
// curl  localhost:9090/1 -XPUT | jq
// PUT with  data
// curl  localhost:9090/1 -XPUT -d '{"id":1, "name": "tea", "description": "hehe"}' | jq
// POST with  data
// curl localhost:9090/1 -X POST -d '{"name": "bubble"}' | jq
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
