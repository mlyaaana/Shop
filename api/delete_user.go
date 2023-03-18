package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	a.userService.Delete(id)

	return c.NoContent(http.StatusOK)
}
