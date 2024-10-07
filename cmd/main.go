package main

import (
	"net/http"
	"rest-api/internal/pkg/centralized"
	"rest-api/internal/pkg/chi"
)

func main() {
	centralized.Centralized()
	routes := chi.NewChiRoutes()
	Listen("localhost:8181", routes)
}

func Listen(addr string, handler http.Handler) {
	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
