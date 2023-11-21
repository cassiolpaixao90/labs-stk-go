package in

import (
	"context"
	"labs-stk-go/internal/application/domain"
)

type GetAllBooksUseCase interface {
	GetAllBooks(ctx context.Context, id string) (*domain.BookDomain, error)
}
