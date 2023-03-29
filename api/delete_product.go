package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDeleteProduct(c echo.Context) error {
	id := c.FormValue("id")

	a.productService.Delete(id)

	return c.NoContent(http.StatusNoContent)
}
