package in

import "labs-stk-go/internal/core/domain"

type ICreate interface {
	Create(user *domain.BookDomain) error
}

type IGetByID interface {
	GetByID(id int) (*domain.BookDomain, error)
}

type IBookServicePort interface {
	ICreate
	IGetByID
}
