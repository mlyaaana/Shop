package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shop/api"
	"shop/repository/credentials"
	productRepo "shop/repository/product"
	userRepo "shop/repository/user"
	"shop/service/auth"
	"shop/service/product"
	"shop/service/user"
)

type App struct {
	echo *echo.Echo
	api  *api.Api
}

func New() *App {
	userRepository := userRepo.NewMapRepository()
	productRepository := productRepo.NewMapRepository()
	credentialsRepository := credentials.NewMapRepository()
	userService := user.NewService(userRepository)
	productService := product.NewService(productRepository)
	authService := auth.NewService(userRepository, credentialsRepository)

	e := echo.New()
	e.Use(middleware.Logger(), middleware.CORS())

	return &App{
		echo: e,
		api:  api.NewApi(userService, productService, authService),
	}
}

func (a *App) Start() {
	a.echo.Logger.Fatal(a.echo.Start(":8080"))
}

func (a *App) InitRoutes() {
	a.echo.POST("/product", a.api.HandleCreateProduct)
}
