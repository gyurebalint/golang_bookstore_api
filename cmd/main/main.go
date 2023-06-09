package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gyurebalint/golang_bookstore_api/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
