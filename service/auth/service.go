package auth

import (
	"shop/domain"
	"shop/repository/credentials"
	"shop/repository/user"
)

type Service struct {
	userRepo        user.Repository
	credentialsRepo credentials.Repository
}

func NewService(userRepo user.Repository, credentialsRepo credentials.Repository) *Service {
	return &Service{
		userRepo:        userRepo,
		credentialsRepo: credentialsRepo,
	}
}

func (s *Service) Register(username, password string) (string, error) {
	u := domain.NewUser(username)
	c := domain.NewCredentials(u.Id, username, password)
	if err := s.userRepo.Create(u); err != nil {
		return "", err
	}
	err := s.credentialsRepo.Create(c)
	if err != nil {
		return "", err
	}
	return u.Id, nil
}

func (s *Service) Auth(username, password string) (string, error) {
	return s.credentialsRepo.Get(username, password)
}
