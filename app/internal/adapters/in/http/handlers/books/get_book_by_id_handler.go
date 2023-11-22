package books

import "github.com/labstack/echo/v4"

type GetBookByIdHandler struct{}

type IGetBookByIdHandler interface {
	GetBookById(c echo.Context) error
}

func NewGetBookByIdIdHandler() IGetBookByIdHandler {
	return &GetBookByIdHandler{}
}

func (b *GetBookByIdHandler) GetBookById(c echo.Context) error {
	return c.JSON(200, struct{}{})
}
