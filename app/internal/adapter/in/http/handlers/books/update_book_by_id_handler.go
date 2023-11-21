package books

import "github.com/labstack/echo/v4"

type UpdateBookByIdHandler struct{}

type IUpdateBookByIdHandler interface {
	UpdateBookById(c echo.Context) error
}

func NewUpdateBookByIdHandler() IUpdateBookByIdHandler {
	return &UpdateBookByIdHandler{}
}

func (b *UpdateBookByIdHandler) UpdateBookById(c echo.Context) error {
	return c.JSON(204, struct{}{})
}
