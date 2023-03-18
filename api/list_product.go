package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListProducts(c echo.Context) error {
	product := a.productService.List()

	return c.JSON(http.StatusOK, product)
}
