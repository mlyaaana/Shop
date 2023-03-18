package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/domain"
)

func (a *Api) HandleCreateProduct(c echo.Context) error {
	body := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{}

	if err := c.Bind(&body); err != nil {
		return err
	}

	product := domain.NewProduct(body.Name, body.Description)
	err := a.productService.Create(product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, product)
}
