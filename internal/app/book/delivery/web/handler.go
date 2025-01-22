package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"rest-api/internal/app/book/usecase"
)

func MakeGetAllBooks(usecase usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resp, err := usecase.Get(ctx)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, resp)
		return
	}
}

func MakeRequestBook(book usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var cmd usecase.BookRequest
		if errDecode := render.DecodeJSON(r.Body, &cmd); errDecode != nil {
			render.Status(r, http.StatusBadRequest)
			return
		}

		err := book.Create(ctx, cmd)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		return
	}
}

func MakeGetBookById(book usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")

		bookResp, err := book.GetById(ctx, id)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.JSON(w, r, bookResp)
		render.Status(r, http.StatusOK)
		return
	}
}
