package application

import (
	"product-api/internal/domain"
	"product-api/internal/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(id string) (*domain.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) GetProductMultiple(ids []string) ([]*domain.Product, error) {
	return s.repo.GetProducts(ids)
}
