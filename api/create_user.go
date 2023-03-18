package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/domain"
)

func (a *Api) HandleCreateUser(c echo.Context) error {
	name := c.FormValue("name")

	user := domain.NewUser(name)
	err := a.userService.Create(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
