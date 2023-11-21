package handlers

import "labs-stk-go/internal/adapter/in/http/handlers/books"

func NewGetBooksInstance() books.IGetBooksHandler {
	return books.NewGetBooksHandler()
}

func NewCreateBookInstance() books.ICreateBookHandler {
	return books.NewCreateBookHandler()
}

func NewGetBookByIdInstance() books.IGetBookByIdHandler {
	return books.NewGetBookByIdIdHandler()
}

func NewUpdateBookByIdInstance() books.IUpdateBookByIdHandler {
	return books.NewUpdateBookByIdHandler()
}

func NewDeleteBookByIdInstance() books.IDeleteBookByIdHandler {
	return books.NewDeleteBookByIdHandler()
}
