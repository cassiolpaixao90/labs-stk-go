package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	echo *echo.Echo
}

func NewServer() *HttpServer {
	return &HttpServer{echo: echo.New()}
}

func (s *HttpServer) Start() {
	// Middleware
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	// Routes
	r := NewHttpRoutes(s.echo)
	r.SetupRouter()

	s.echo.Logger.Fatal(s.echo.Start(":9092"))
}
