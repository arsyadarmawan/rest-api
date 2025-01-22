package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
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
		3,
		time.Second*30,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, map[string]interface{}{"error": "rate limit exceeded"})
			render.Status(r, http.StatusTooManyRequests)
		}),
	)).Group(func(r chi.Router) {
		r.Post("/book", MakeRequestBook(b.Opts.Book))
		r.Get("/books", MakeGetAllBooks(b.Opts.Book))
		r.Get("/book/{id}", MakeGetBookById(b.Opts.Book))
	})
	return nil
}
