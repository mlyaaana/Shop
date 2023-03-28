package product

import (
	"shop/domain"
	"shop/repository/product"
)

type Service struct {
	repo product.Repository
}

func NewService(repo product.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(product *domain.Product) error {
	return s.repo.Create(product)
}

func (s *Service) Get(id string) (*domain.Product, error) {
	return s.repo.Get(id)
}

func (s *Service) List() ([]*domain.Product, error) {
	return s.repo.List()
}

func (s *Service) Delete(id string) {
	s.repo.Delete(id)
}
