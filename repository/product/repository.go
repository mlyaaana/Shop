package product

import "Shop/domain"

type Repository interface {
	Create(product *domain.Product)
	Get(id string) (*domain.Product, error)
	List() []*domain.Product
	Delete(id string)
}
