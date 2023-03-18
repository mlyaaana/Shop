package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleGetProduct(c echo.Context) error {
	id := c.FormValue("id")

	product, err := a.productService.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}
