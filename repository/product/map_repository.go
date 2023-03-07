package product

import (
	"Shop/domain"
	"errors"
)

type MapRepository struct {
	storage map[string]*domain.Product
}

func NewMapRepository() *MapRepository {
	return &MapRepository{
		storage: make(map[string]*domain.Product),
	}
}

func (m *MapRepository) Create(product *domain.Product) {
	m.storage[product.Id] = product
}

func (m *MapRepository) Get(id string) (*domain.Product, error) {
	product, ok := m.storage[id]
	if !ok {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (m *MapRepository) List() []*domain.Product {
	res := make([]*domain.Product, 0)
	for _, v := range m.storage {
		res = append(res, v)
	}

	return res
}

func (m *MapRepository) Delete(id string) {
	delete(m.storage, id)
}
