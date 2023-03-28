package product

import "shop/domain"

type Repository interface {
	Create(product *domain.Product) error
	Get(id string) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(id string)
}
