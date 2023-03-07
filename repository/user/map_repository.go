package user

import (
	"Shop/domain"
	"errors"
)

type MapRepository struct {
	storage map[string]*domain.User
}

func NewMapRepository() *MapRepository {
	return &MapRepository{
		storage: make(map[string]*domain.User),
	}
}

func (m *MapRepository) Create(user *domain.User) error {
	for _, v := range m.storage {
		if v.Name == user.Name {
			return errors.New("username is already taken")
		}
	}
	m.storage[user.Id] = user

	return nil
}

func (m *MapRepository) Get(id string) (*domain.User, error) {
	user, ok := m.storage[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (m *MapRepository) List() []*domain.User {
	res := make([]*domain.User, 0)

	for _, v := range m.storage {
		res = append(res, v)
	}

	return res
}

func (m *MapRepository) Delete(id string) {
	delete(m.storage, id)
}
