package services

import (
	"context"
	"labs-stk-go/internal/application/domain"
	"labs-stk-go/internal/application/ports/in"
	"labs-stk-go/internal/application/ports/out"
)

type GetAllBooksService struct {
	GetAllBooksClientPort out.GetAllBooksClientPort
}

func NewGetAllBooksService(
	getAllBooksByIdPort out.GetAllBooksClientPort,
) in.GetAllBooksUseCase {
	return &GetAllBooksService{getAllBooksByIdPort}
}

func (g *GetAllBooksService) GetAllBooks(ctx context.Context, id string) (*domain.BookDomain, error) {
	book, err := g.GetAllBooksClientPort.GetAllBooks(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
