package api

import (
	"Shop/service/product"
	"Shop/service/user"
)

type Api struct {
	userService    *user.Service
	productService *product.Service
}

func NewApi(userService *user.Service, productService *product.Service) *Api {
	return &Api{
		userService:    userService,
		productService: productService,
	}
}
