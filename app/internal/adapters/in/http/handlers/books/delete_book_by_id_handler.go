package books

import "github.com/labstack/echo/v4"

type DeleteBookByIdHandler struct{}

type IDeleteBookByIdHandler interface {
	DeleteBookById(c echo.Context) error
}

func NewDeleteBookByIdHandler() IDeleteBookByIdHandler {
	return &DeleteBookByIdHandler{}
}

func (b *DeleteBookByIdHandler) DeleteBookById(c echo.Context) error {
	return c.JSON(204, struct{}{})
}
