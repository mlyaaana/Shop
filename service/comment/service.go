package comment

import (
	"shop/domain"
	"shop/repository/comment"
)

type Service struct {
	repo comment.Repository
}

func NewService(repo comment.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(comment *domain.Comment) error {
	return s.repo.Create(comment)
}

func (s *Service) List(productId string) ([]*domain.Comment, error) {
	return s.repo.List(productId)
}

func (s *Service) Delete(id, userId string) {
	s.repo.Delete(id, userId)
}

func (s *Service) AdminDelete(id string) {
	s.repo.AdminDelete(id)
}
