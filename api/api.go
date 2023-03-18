package api

import (
	"shop/service/auth"
	"shop/service/product"
	"shop/service/user"
)

type Api struct {
	userService    *user.Service
	productService *product.Service
	authService    *auth.Service
}

func NewApi(userService *user.Service, productService *product.Service, authService *auth.Service) *Api {
	return &Api{
		userService:    userService,
		productService: productService,
		authService:    authService,
	}
}
