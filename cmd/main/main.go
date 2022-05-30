package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jayeshsadhwani99/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", nil))
}