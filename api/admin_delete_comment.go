package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Api) HandleAdminDeleteComment(c echo.Context) error {
	id := c.FormValue("id")

	a.commentService.AdminDelete(id)

	return c.NoContent(http.StatusNoContent)
}
