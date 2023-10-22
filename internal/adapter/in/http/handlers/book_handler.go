package handlers

import (
	"labs-stk-go/internal/core/domain"
	"labs-stk-go/internal/core/ports/in"
)

type IBookHandler interface {
	GetByID(id int) (*domain.BookDomain, error)
	Create(user *domain.BookDomain) error
}

type BookHandler struct {
	in.IBookServicePort
}

func NewBookHandler() IBookHandler {
	return &BookHandler{}
}

func (b *BookHandler) GetByID(id int) (*domain.BookDomain, error) {
	_, err := b.GetByID(id)
	if err != nil {
		return nil, err
	}
	panic("implement me")
}

func (b *BookHandler) Create(user *domain.BookDomain) error {
	err := b.Create(user)
	if err != nil {
		return err
	}
	panic("implement me")
}
