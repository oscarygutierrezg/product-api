package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewInMemoryProductRepository_CreatesFiveProducts
func TestNewInMemoryProductRepository_CreatesFiveProducts(t *testing.T) {
	// Given
	repo := NewInMemoryProductRepository()

	// When
	actualCount := len(repo.products)

	// Then
	assert.Equal(t, 5, actualCount, "Expected 5 products")
}

// TestGetProductByID_ValidID_ReturnsProduct
func TestGetProductByID_ValidID_ReturnsProduct(t *testing.T) {
	// Given
	repo := NewInMemoryProductRepository()
	productID := "1fafb319-67d9-4cc3-9182-a1099c0919cb"

	// When
	product, err := repo.GetProductByID(productID)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, product, "Expected a product, got nil")
	assert.Equal(t, productID, product.ID, "Expected product ID to match")
}

// TestGetProductByID_InvalidID_ReturnsError
func TestGetProductByID_InvalidID_ReturnsError(t *testing.T) {
	// Given
	repo := NewInMemoryProductRepository()
	productID := "999" // ID que no existe

	// When
	product, err := repo.GetProductByID(productID)

	// Then
	assert.Error(t, err, "Expected an error, got none")
	assert.Nil(t, product, "Expected nil product, got a product")
}

// TestGetProducts_ValidIDs_ReturnsProducts
func TestGetProducts_ValidIDs_ReturnsProducts(t *testing.T) {
	// Given
	repo := NewInMemoryProductRepository()
	productIDs := []string{
		"1fafb319-67d9-4cc3-9182-a1099c0919cb",
		"9023995c-72af-4170-b449-482ec5df146b",
	}

	// When
	products, err := repo.GetProducts(productIDs)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, products, "Expected products, got nil")
	assert.Equal(t, len(productIDs), len(products), "Expected number of products to match the number of IDs")
	for _, id := range productIDs {
		found := false
		for _, product := range products {
			if product.ID == id {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected to find product with ID: "+id)
	}
}

// TestGetProducts_InvalidID_ReturnsError
func TestGetProducts_InvalidID_ReturnsError(t *testing.T) {
	// Given
	repo := NewInMemoryProductRepository()
	productIDs := []string{
		"999", // ID que no existe
	}

	// When
	products, err := repo.GetProducts(productIDs)

	// Then
	assert.Error(t, err, "Expected an error, got none")
	assert.Nil(t, products, "Expected nil products, got some products")
}
