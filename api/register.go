package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleRegister(c echo.Context) error {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&body); err != nil {
		return err
	}

	userId, err := a.authService.Register(body.Username, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userId)
}
