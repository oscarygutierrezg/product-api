package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-api/internal/adapters/api/v1"
	"product-api/internal/adapters/repository"
	"product-api/internal/application"
)

func main() {

	repo := repository.NewInMemoryProductRepository()
	productService := application.NewProductService(repo)
	productHandler := v1.NewProductHandler(productService)

	r := mux.NewRouter()

	r.HandleFunc("/v1/products/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/v1/products/multiple", productHandler.GetProductMultiple).Methods("POST")

	log.Println("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
