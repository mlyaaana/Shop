package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListUsers(c echo.Context) error {
	user := a.userService.List()

	return c.JSON(http.StatusOK, user)
}
