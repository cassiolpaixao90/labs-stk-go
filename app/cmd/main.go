package main

import (
	"labs-stk-go/internal/adapter/in/http"
)

func main() {
	server := http.NewServer()
	server.Start()
}
