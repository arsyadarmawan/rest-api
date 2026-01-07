package web

import (
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func MakeGetAllBooks(usc usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cmd := usecase.BookGetter{
			Limit:  r.URL.Query().Get("limit"),
			Offset: r.URL.Query().Get("offset"),
		}
		resp, err := usc.Get(ctx, cmd)
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
		if err := r.ParseMultipartForm(10 << 20); err != nil { // Limit to 10MB
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "failed to parse form-data"})
			return
		}

		cmd.Name = r.FormValue("name")
		cmd.Description = r.FormValue("description")
		cmd.ReleaseDate = r.FormValue("releaseDate")
		cmd.PublishedBy = r.FormValue("publishedBy")
		cmd.Author = r.FormValue("author")

		file, _, err := r.FormFile("image")

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "failed to retrieve image file"})
			return
		}
		defer file.Close()
		if errDecode := render.DecodeJSON(r.Body, &cmd); errDecode != nil {
			render.Status(r, http.StatusBadRequest)
			return
		}

		errCreate := book.Create(ctx, cmd)
		if errCreate != nil {
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

func MakeDeleteBookById(book usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")

		err := book.Delete(ctx, id)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		return
	}
}
