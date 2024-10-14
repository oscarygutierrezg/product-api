package ports

import "product-api/internal/domain"

//go:generate mockery --name=ProductRepository --output=. --outpkg=repository --filename=mock_product_repository.go
type ProductRepository interface {
	GetProductByID(id string) (*domain.Product, error)

	GetProducts(ids []string) ([]*domain.Product, error)
}
