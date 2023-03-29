package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleListComment(c echo.Context) error {
	productId := c.FormValue("productId")

	comment, err := a.commentService.List(productId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}
