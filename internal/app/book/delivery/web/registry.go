package web

import (
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
	"github.com/go-chi/chi"
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

//func (b BookRegistry) RegisterRoutesTo(r *chi.Mux) *chi.Mux {
//	r.With(httprate.Limit(
//		3,
//		time.Second*30,
//		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
//			render.Status(r, http.StatusTooManyRequests)
//			//render.JSON(w, r, map[string]interface{}{"error": "rate limit exceeded"})
//			return
//		}),
//	)).Group(func(r chi.Router) {
//		r.Post("/book", MakeRequestBook(b.Opts.Book))
//		r.Get("/books", MakeGetAllBooks(b.Opts.Book))
//		r.Get("/book/{id}", MakeGetBookById(b.Opts.Book))
//		r.Delete("/book/{id}", MakeDeleteBookById(b.Opts.Book))
//	})
//	return nil
//}

func (b BookRegistry) RegisterRoutesTo(r *chi.Mux) *chi.Mux {
	//r.With(httprate.Limit(
	//	3,
	//	time.Second*30,
	//	httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
	//		render.Status(r, http.StatusTooManyRequests)
	//		//render.JSON(w, r, map[string]interface{}{"error": "rate limit exceeded"})
	//		return
	//	}),
	//)).Group(func(r chi.Router) {
	r.Post("/book", MakeRequestBook(b.Opts.Book))
	r.Get("/books", MakeGetAllBooks(b.Opts.Book))
	r.Get("/book/{id}", MakeGetBookById(b.Opts.Book))
	r.Delete("/book/{id}", MakeDeleteBookById(b.Opts.Book))
	//})
	return nil
}
