package http

import (
	"github.com/labstack/echo/v4"
	"labs-stk-go/internal/adapter/in/http/handlers"
)

type Routes struct {
	echo *echo.Echo
}

func NewHttpRoutes(e *echo.Echo) *Routes {
	return &Routes{echo: e}
}

func (r *Routes) SetupRouter() {
	r.echo.GET("/books", handlers.NewGetBooksInstance().GetBooks)
	r.echo.POST("/books", handlers.NewCreateBookInstance().CreateBook)
	r.echo.GET("/books/:id", handlers.NewGetBookByIdInstance().GetBookById)
	r.echo.PUT("/books/:id", handlers.NewUpdateBookByIdInstance().UpdateBookById)
	r.echo.DELETE("/books/:id", handlers.NewDeleteBookByIdInstance().DeleteBookById)
}
