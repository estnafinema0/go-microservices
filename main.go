package main

import (
	"log"
	"net/http"
	"os"

	"github.com/estnafinema0/go-microservices/handlers"
)

//go run main.go
//curl -v -d "Hello World" localhost:9090/...

func main() {

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Goodbye World")
	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
