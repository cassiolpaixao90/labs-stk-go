package handlers

import (
	"labs-stk-go-hexagonal/internal/core/ports/in"
	"labs-stk-go-hexagonal/internal/domain"
)

type IBookHandler interface {
	GetByID(id int) (*domain.User, error)
	Create(user *domain.User) error
}

type BookHandler struct {
	in.IBookServicePort
}

func NewBookHandler() IBookHandler {
	return &BookHandler{}
}

func (b *BookHandler) GetByID(id int) (*domain.User, error) {
	_, err := b.GetByID(id)
	if err != nil {
		return nil, err
	}
	panic("implement me")
}

func (b *BookHandler) Create(user *domain.User) error {
	err := b.Create(user)
	if err != nil {
		return err
	}
	panic("implement me")
}
