package application

import (
	"errors"
	"product-api/internal/domain"
	"product-api/internal/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetProduct_ValidID_ReturnsProduct prueba el caso en que se solicita un producto válido.
func TestGetProduct_ValidID_ReturnsProduct(t *testing.T) {
	// Given
	mockRepo := &ports.MockProductRepository{}

	productID := "1fafb319-67d9-4cc3-9182-a1099c0919cb"
	expectedProduct := &domain.Product{
		ID:          productID,
		Name:        "Laptop",
		Description: "High-end device",
		Price:       1500,
	}

	mockRepo.On("GetProductByID", productID).Return(expectedProduct, nil)

	productService := NewProductService(mockRepo)

	// When
	product, err := productService.GetProduct(productID)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, product, "Expected product, got nil")
	assert.Equal(t, expectedProduct.ID, product.ID, "Expected product ID to match")
	mockRepo.AssertExpectations(t)
}

// TestGetProduct_InvalidID_ReturnsError prueba el caso en que se solicita un producto no válido.
func TestGetProduct_InvalidID_ReturnsError(t *testing.T) {
	// Given
	mockRepo := &ports.MockProductRepository{}

	invalidID := "1"

	mockRepo.On("GetProductByID", invalidID).Return(nil, errors.New("product not found"))

	productService := NewProductService(mockRepo)

	// When
	product, err := productService.GetProduct(invalidID)

	// Then
	assert.Error(t, err, "Expected an error")
	assert.Nil(t, product, "Expected nil product")
	assert.Equal(t, "product not found", err.Error(), "Expected 'product not found' error")
	mockRepo.AssertExpectations(t)
}

// TestGetProductMultiple_ValidIDs_ReturnsProducts prueba el caso en que se solicitan múltiples productos válidos.
func TestGetProductMultiple_ValidIDs_ReturnsProducts(t *testing.T) {
	// Given
	mockRepo := &ports.MockProductRepository{}

	productIDs := []string{
		"1fafb319-67d9-4cc3-9182-a1099c0919cb",
		"9023995c-72af-4170-b449-482ec5df146b",
	}
	expectedProducts := []*domain.Product{
		{
			ID:          "1fafb319-67d9-4cc3-9182-a1099c0919cb",
			Name:        "Laptop",
			Description: "High-end device",
			Price:       1500,
		},
		{
			ID:          "9023995c-72af-4170-b449-482ec5df146b",
			Name:        "Smartphone",
			Description: "Flagship phone",
			Price:       1000,
		},
	}

	mockRepo.On("GetProducts", productIDs).Return(expectedProducts, nil)

	productService := NewProductService(mockRepo)

	// When
	products, err := productService.GetProductMultiple(productIDs)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, products, "Expected products, got nil")
	assert.Equal(t, len(expectedProducts), len(products), "Expected number of products to match")

	for i, product := range products {
		assert.Equal(t, expectedProducts[i].ID, product.ID, "Expected product ID to match")
	}

	mockRepo.AssertExpectations(t)
}

// TestGetProductMultiple_InvalidID_ReturnsError prueba el caso en que se solicita un ID inválido en la lista.
func TestGetProductMultiple_InvalidID_ReturnsError(t *testing.T) {
	// Given
	mockRepo := &ports.MockProductRepository{}

	productIDs := []string{
		"invalid-id", // ID inválido
	}

	mockRepo.On("GetProducts", productIDs).Return(nil, errors.New("one or more products not found"))

	productService := NewProductService(mockRepo)

	// When
	products, err := productService.GetProductMultiple(productIDs)

	// Then
	assert.Error(t, err, "Expected an error")
	assert.Nil(t, products, "Expected nil products")
	assert.Equal(t, "one or more products not found", err.Error(), "Expected 'one or more products not found' error")
	mockRepo.AssertExpectations(t)
}
