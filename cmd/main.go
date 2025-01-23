package main

import (
	"github.com/arsyadarmawan/rest-api/internal/pkg/asynq"
	"github.com/arsyadarmawan/rest-api/internal/pkg/centralized"
	"github.com/arsyadarmawan/rest-api/internal/pkg/chi"
	"net/http"
)

func main() {
	centralized.Centralized()
	go asynq.InitServeMuxAsynq()
	routes := chi.NewChiRoutes()
	Listen("localhost:8181", routes)
}

func Listen(addr string, handler http.Handler) {
	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
