package books

import "github.com/labstack/echo/v4"

type CreateBookHandler struct{}

type ICreateBookHandler interface {
	CreateBook(c echo.Context) error
}

func NewCreateBookHandler() ICreateBookHandler {
	return &CreateBookHandler{}
}

func (b *CreateBookHandler) CreateBook(c echo.Context) error {
	return c.JSON(201, struct{}{})
}
