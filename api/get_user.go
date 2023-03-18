package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleGetUser(c echo.Context) error {
	id := c.FormValue("id")

	user, err := a.userService.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
