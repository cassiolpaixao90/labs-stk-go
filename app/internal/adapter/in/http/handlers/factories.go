package handlers

import (
	"labs-stk-go/internal/adapter/in/http/handlers/books"
	book_service "labs-stk-go/internal/adapter/out/http/client/book-service"
	"labs-stk-go/internal/application/services"
)

func NewGetBooksInstance() books.IGetAllBooksHandler {
	client := book_service.NewGetAllBooksClient()
	service := services.NewGetAllBooksService(client)
	return books.NewGetAllBooksHandler(service)
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
