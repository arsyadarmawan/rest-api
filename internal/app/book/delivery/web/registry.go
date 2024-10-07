package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	"net/http"
	"rest-api/internal/app/book/usecase"
	"time"
)

type BookRegistry struct {
	Opts BookRegistryOpts
}

type BookRegistryOpts struct {
	Book usecase.Book
}

func NewBookRegistry(opts BookRegistryOpts) *BookRegistry {
	return &BookRegistry{
		Opts: opts,
	}
}

func (b BookRegistry) RegisterRoutesTo(r *chi.Mux) *chi.Mux {
	r.With(httprate.Limit(
		1,
		time.Minute,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, `{"error": "Rate-limited. Please, slow down."}`, http.StatusTooManyRequests)
		}),
	)).Group(func(r chi.Router) {
		// POST route for creating a new book
		r.Post("/book", MakeRequestBook(b.Opts.Book))
	})
	r.Get("/books/{id}", MakeGetAllBooks(b.Opts.Book))
	return nil
}
