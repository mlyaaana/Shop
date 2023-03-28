package user

import (
	"shop/domain"
	"shop/repository/user"
)

type Service struct {
	repo user.Repository
}

func NewService(repo user.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *Service) Get(id string) (*domain.User, error) {
	return s.repo.Get(id)
}

func (s *Service) List() ([]*domain.User, error) {
	return s.repo.List()
}

func (s *Service) Delete(id string) {
	s.repo.Delete(id)
}
