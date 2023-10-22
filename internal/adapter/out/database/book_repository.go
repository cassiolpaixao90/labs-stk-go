package database

import (
	"context"
	"labs-stk-go/internal/core/domain"
	"labs-stk-go/pkg/mongo"
)

type IBookRepository interface {
	Create(*domain.BookDomain) (*domain.BookDomain, error)
}

type BookRepository struct {
	DB mongo.IMongoRepository[BookSchema, domain.BookDomain]
}

func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

func (br *BookRepository) Create(bookDomain *domain.BookDomain) (*domain.BookDomain, error) {
	_, err := br.DB.InsertOne(&BookSchema{ID: "id", Name: ""}, context.TODO())
	if err != nil {
		return nil, err
	}
	return bookDomain, nil
}
