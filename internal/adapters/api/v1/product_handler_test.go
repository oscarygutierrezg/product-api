package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"product-api/internal/adapters/repository"
	"product-api/internal/application"
	"product-api/internal/domain"
	"testing"
)

func TestGetProduct_ValidID_ReturnsProduct(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryProductRepository()
	service := application.NewProductService(mockRepository)

	handler := NewProductHandler(service)

	req, err := http.NewRequest("GET", "/v1/products/1fafb319-67d9-4cc3-9182-a1099c0919cb", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/products/{id}", handler.GetProduct)

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusOK, rr.Code)

	var product domain.Product
	if err := json.NewDecoder(rr.Body).Decode(&product); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	assert.NotNil(t, &product)
	assert.NotNil(t, &product.ID)
	assert.NotNil(t, &product.Name)
	assert.NotNil(t, &product.Description)
	assert.NotNil(t, &product.Price)

}

func TestGetProduct_NotFound_Returns404(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryProductRepository()
	service := application.NewProductService(mockRepository)

	handler := NewProductHandler(service)

	// Crear una solicitud para un producto que no existe
	req, err := http.NewRequest("GET", "/v1/products/non-existent-id", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/products/{id}", handler.GetProduct).Methods("GET")

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "product not found\n", rr.Body.String())
}
func TestGetProductMultiple_ValidIDs_ReturnsProducts(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryProductRepository()
	service := application.NewProductService(mockRepository)

	handler := NewProductHandler(service)

	ids := []string{
		"1fafb319-67d9-4cc3-9182-a1099c0919cb",
		"9023995c-72af-4170-b449-482ec5df146b",
	}

	body, err := json.Marshal(ids)
	if err != nil {
		t.Fatalf("could not marshal ids: %v", err)
	}
	req, err := http.NewRequest("POST", "/v1/products/multiple", bytes.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/products/multiple", handler.GetProductMultiple).Methods("POST")

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusOK, rr.Code)

	var products []*domain.Product
	if err := json.NewDecoder(rr.Body).Decode(&products); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	assert.NotNil(t, products)
	assert.Equal(t, 2, len(products), "Expected two products")
	for _, product := range products {
		assert.NotEmpty(t, product.ID)
		assert.NotEmpty(t, product.Name)
		assert.NotEmpty(t, product.Description)
		assert.NotZero(t, product.Price)
	}
}

func TestGetProductMultiple_ValidIDs_ReturnsErrorInvalid_Request_Payload(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryProductRepository()
	service := application.NewProductService(mockRepository)

	handler := NewProductHandler(service)

	req, err := http.NewRequest("POST", "/v1/products/multiple", httptest.NewRecorder().Body)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/products/multiple", handler.GetProductMultiple).Methods("POST")

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetProductMultiple_NotFound_Returns404(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryProductRepository()
	service := application.NewProductService(mockRepository)

	handler := NewProductHandler(service)

	// Crear una lista de IDs donde al menos uno de los productos no existe
	ids := []string{"non-existent-id-1", "non-existent-id-2"}

	// Convertimos los IDs a JSON y creamos un cuerpo de solicitud con `bytes.NewReader`
	body, err := json.Marshal(ids)
	if err != nil {
		t.Fatalf("could not marshal ids: %v", err)
	}

	req, err := http.NewRequest("POST", "/v1/products/multiple", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/products/multiple", handler.GetProductMultiple).Methods("POST")

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "product not found for id: non-existent-id-1\n", rr.Body.String())
}
