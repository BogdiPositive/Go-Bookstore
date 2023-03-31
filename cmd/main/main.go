package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bogdipositive/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server listening on http://localhost:8080\n")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
