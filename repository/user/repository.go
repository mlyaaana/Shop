package user

import "shop/domain"

type Repository interface {
	Create(user *domain.User) error
	Get(id string) (*domain.User, error)
	List() []*domain.User
	Delete(id string)
}
