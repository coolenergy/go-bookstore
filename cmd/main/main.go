package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Cerebrovinny/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println("Server started on port 9010")
}
