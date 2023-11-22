package main

import (
	"labs-stk-go/internal/adapters/in/http"
)

func main() {
	server := http.NewServer()
	server.Start()
}
