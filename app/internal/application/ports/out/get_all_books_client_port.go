package out

import (
	"context"
	"labs-stk-go/internal/application/domain"
)

type GetAllBooksClientPort interface {
	GetAllBooks(ctx context.Context, id string) (*domain.BookDomain, error)
}
