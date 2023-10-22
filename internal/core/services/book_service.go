package services

import (
	"labs-stk-go-hexagonal/internal/core/domain"
	"labs-stk-go-hexagonal/internal/core/ports/in"
)

type BookService struct {
}

func NewBookService() in.IBookServicePort {
	return &BookService{}
}

func (b *BookService) Create(*domain.BookDomain) error {
	panic("implement me")
}

func (b *BookService) GetByID(id int) (*domain.BookDomain, error) {
	panic("implement me")
}
