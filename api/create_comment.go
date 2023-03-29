package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shop/domain"
)

func (a *Api) HandleCreateComment(c echo.Context) error {
	body := struct {
		UserId    string `json:"userId"`
		ProductId string `json:"productId"`
		Mention   string `json:"mention"`
	}{}

	if err := c.Bind(&body); err != nil {
		return err
	}

	comment := domain.NewComment(body.UserId, body.ProductId, body.Mention)
	err := a.commentService.Create(comment)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, comment)
}
