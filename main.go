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
)

//go run main.go
//curl -v -d "Hello World" localhost:9090/...
//curl -v -d "Nice" localhost:9090 | jq
// another method that migth not be allowed
// curl  localhost:9090 -XDELETE -v | jq
//POST with  json
//curl  localhost:9090 -d '{"id":1, "name": "tea", "description": "hehe"}' | jq
//PUT with  json
//curl  localhost:9090/1 -XPUT | jq
//PUT with  data
//curl  localhost:9090/1 -XPUT -d '{"id":1, "name": "tea", "description": "hehe"}' | jq

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	//sm.Handle("/goodbye", gh)

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
