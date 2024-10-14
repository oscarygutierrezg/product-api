package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"product-api/internal/application"
)

type ProductHandler struct {
	service *application.ProductService
}

func NewProductHandler(service *application.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := vars["id"]

	product, err := h.service.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProductMultiple(w http.ResponseWriter, r *http.Request) {
	var ids []string

	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	products, err := h.service.GetProductMultiple(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
