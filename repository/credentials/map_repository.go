package credentials

import (
	"errors"
	"shop/domain"
)

type MapRepository struct {
	storage map[string]*domain.Credentials
}

func NewMapRepository() *MapRepository {
	return &MapRepository{
		storage: make(map[string]*domain.Credentials),
	}
}

func (m *MapRepository) Create(credentials *domain.Credentials) error {
	m.storage[credentials.Username] = credentials
	return nil
}

func (m *MapRepository) Get(username, password string) (string, error) {
	if creds, ok := m.storage[username]; ok && creds.Password == password {
		return creds.UserId, nil
	}

	return "", errors.New("incorrect username or password")
}
