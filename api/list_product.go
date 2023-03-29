package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListProducts(c echo.Context) error {
	product, err := a.productService.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}
