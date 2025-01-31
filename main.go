package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(wr, "Oops", http.StatusBadRequest)
			// wr.WriteHeader(http.StatusBadRequest)
			// wr.Write([]byte("Oops"))
			return
		}
		//log.Printf("Data %s\n", d)
		fmt.Fprintf(wr, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
