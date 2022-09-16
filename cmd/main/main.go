package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imvedanshmehra/books-management/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
