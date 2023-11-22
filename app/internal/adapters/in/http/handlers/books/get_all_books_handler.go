package books

import (
	"context"
	"github.com/labstack/echo/v4"
	"labs-stk-go/internal/application/ports/in"
)

type GetAllBooksHandler struct {
	in.GetAllBooksUseCase
}

type IGetAllBooksHandler interface {
	GetAllBooks(c echo.Context) error
}

func NewGetAllBooksHandler(getAllBooksUseCase in.GetAllBooksUseCase) IGetAllBooksHandler {
	return &GetAllBooksHandler{getAllBooksUseCase}
}

func (b *GetAllBooksHandler) GetAllBooks(c echo.Context) error {
	books, err := b.GetAllBooksUseCase.GetAllBooks(context.Background(), "")
	if err != nil {
		return err
	}
	return c.JSON(200, books)
}
