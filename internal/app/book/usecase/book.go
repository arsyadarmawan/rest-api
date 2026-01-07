package usecase

import (
	"context"
	error2 "github.com/arsyadarmawan/rest-api/internal/pkg/error"
)

//go:generate mockgen -source=book.go -destination=usecasemock/book_mock.go -package=usecasemock
type Book interface {
	Get(ctx context.Context, book BookGetter) ([]BookResponse, *error2.AppError)
	GetById(ctx context.Context, id string) (BookResponse, *error2.AppError)
	Create(ctx context.Context, cmd BookRequest) *error2.AppError
	Delete(ctx context.Context, id string) *error2.AppError
	Update(ctx context.Context, id string) *error2.AppError
}
