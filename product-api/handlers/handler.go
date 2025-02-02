package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(wr, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(wr, "Hello %s\n", d)
}
