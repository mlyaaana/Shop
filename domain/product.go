package domain

import (
	"github.com/google/uuid"
)

type Product struct {
	Id          string
	Name        string
	Description string
}

func NewProduct(name, description string) *Product {
	return &Product{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
