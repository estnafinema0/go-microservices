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
	"github.com/nicholasjackson/env"
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

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind Address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
