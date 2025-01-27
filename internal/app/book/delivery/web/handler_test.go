package web

import (
	"context"
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockBookUseCase struct{}

func (m *mockBookUseCase) Get(ctx context.Context) ([]usecase.BookResponse, error) {
	return []usecase.BookResponse{}, nil
}

func (m *mockBookUseCase) GetById(ctx context.Context, id string) (usecase.BookResponse, error) {
	return usecase.BookResponse{}, nil
}

func (m *mockBookUseCase) Create(ctx context.Context, cmd usecase.BookRequest) error {
	return nil
}

func (m *mockBookUseCase) Delete(ctx context.Context, id string) error {
	return nil
}

func (m *mockBookUseCase) Update(ctx context.Context, id string) error {
	return nil
}

func TestMakeGetAllBooks(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{"successful GET all books", http.StatusOK},
		{"successful GET all books", http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/books", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := MakeGetAllBooks(&mockBookUseCase{})
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, test.expectedStatus)
			}
		})
	}
}
