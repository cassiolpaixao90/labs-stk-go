package books

import "github.com/labstack/echo/v4"

type GetBooksHandler struct{}

type IGetBooksHandler interface {
	GetBooks(c echo.Context) error
}

func NewGetBooksHandler() IGetBooksHandler {
	return &GetBooksHandler{}
}

func (b *GetBooksHandler) GetBooks(c echo.Context) error {
	return c.JSON(200, struct{}{})
}
