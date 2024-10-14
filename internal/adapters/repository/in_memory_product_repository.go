package repository

import (
	"errors"
	"math/rand"
	"product-api/internal/domain"
	"time"
)

type InMemoryProductRepository struct {
	products map[string]*domain.Product
}

var productNames = []string{
	"Laptop", "Smartphone", "Tablet", "Smartwatch", "Headphones",
}

var productDescriptions = []string{
	"High-end device", "Budget-friendly", "Latest model", "Lightweight and portable", "Top performance",
}

var productIds = []string{
	"1fafb319-67d9-4cc3-9182-a1099c0919cb",
	"9023995c-72af-4170-b449-482ec5df146b",
	"1d1c3508-107d-469e-a2ac-875dc1533240",
	"193201ef-e120-4f38-9259-04557b04d34a",
	"862f3850-bf60-49ca-a003-5a20dfa96913",
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	rand.Seed(time.Now().UnixNano())
	products := make(map[string]*domain.Product)

	for i := 0; i < 5; i++ {
		products[productIds[i]] = &domain.Product{
			ID:          productIds[i],
			Name:        randomName(),
			Description: randomDescription(),
			Price:       randomPrice(),
		}
	}

	return &InMemoryProductRepository{
		products: products,
	}
}

func randomName() string {
	return productNames[rand.Intn(len(productNames))]
}

func randomDescription() string {
	return productDescriptions[rand.Intn(len(productDescriptions))]
}

func randomPrice() float64 {
	return 50 + rand.Float64()*(2000-50)
}

func (r *InMemoryProductRepository) GetProductByID(id string) (*domain.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *InMemoryProductRepository) GetProducts(ids []string) ([]*domain.Product, error) {
	var foundProducts []*domain.Product

	for _, id := range ids {
		product, exists := r.products[id]
		if !exists {
			return nil, errors.New("product not found for id: " + id)
		}
		foundProducts = append(foundProducts, product)
	}

	return foundProducts, nil
}
