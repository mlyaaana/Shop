package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleDeleteComment(c echo.Context) error {
	id := c.FormValue("id")
	userId := c.FormValue("userId")

	a.commentService.Delete(id, userId)

	return c.NoContent(http.StatusNoContent)
}
