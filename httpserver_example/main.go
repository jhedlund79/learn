package main

import (
	"log"
	"net/http"

	. "./myserver"
)

const addr = "localhost:12345"

func main() {
	mux := http.NewServeMux()
	handler := &MyHandler{}
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/contacts", handler)
	mux.Handle("/products", handler)
	log.Printf("Now listening on %s...\n", addr)
	server := http.Server{Handler: mux, Addr: addr}
	log.Fatal(server.ListenAndServe())
}
