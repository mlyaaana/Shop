package app

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"shop/api"
	"shop/database"
	"shop/repository/credentials"
	productrepo "shop/repository/product"
	userrepo "shop/repository/user"
	"shop/service/auth"
	"shop/service/product"
	"shop/service/user"
)

type App struct {
	echo *echo.Echo
	api  *api.Api
}

func New() *App {
	db, err := database.New()
	if err != nil {
		panic(err)
	}
	userRepository := userrepo.NewDBRepository(db)
	userService := user.NewService(userRepository)
	productService := product.NewService(productrepo.NewDBRepository(db))
	authService := auth.NewService(userRepository, credentials.NewDBRepository(db))
	//userRepository := u.NewDBRepository(db)
	//userService := user.NewService(userRepository)
	//productService := product.NewService(productrepo.NewDBRepository(db))
	//authService := auth.NewService(credentials.NewDBRepository(db), userRepository)

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
	e := a.echo
	r := a.echo.Group("") //требует авторизацию
	secret := os.Getenv("SECRET_KEY")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(api.JwtCustomClaims)
		},
		SigningKey: []byte(secret),
	}
	r.Use(echojwt.WithConfig(config))
	r.POST("/product", a.api.HandleCreateProduct)
	r.POST("/user", a.api.HandleCreateUser)
	r.GET("/users", a.api.HandleListUsers)
	r.GET("/products", a.api.HandleListProducts)
	r.GET("/user", a.api.HandleGetUser)
	r.GET("/product", a.api.HandleGetProduct)
	e.POST("/login", a.api.HandleLogin)
	e.POST("/register", a.api.HandleRegister)
	r.DELETE("/user", a.api.HandleDeleteUser)
	r.DELETE("/product", a.api.HandleDeleteProduct)
}
