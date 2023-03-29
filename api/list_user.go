package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListUsers(c echo.Context) error {
	user, err := a.userService.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
