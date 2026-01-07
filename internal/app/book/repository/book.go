package repository

import (
	"context"
	"github.com/arsyadarmawan/rest-api/internal/app/ent"
	error2 "github.com/arsyadarmawan/rest-api/internal/pkg/error"
)

//go:generate mockgen -source=book.go -destination=repositorymock/book_mock.go -package=repositorymock
type BookRepository interface {
	Get(ctx context.Context, getter ent.BookGetter) ([]*ent.Book, *error2.AppError)
	Create(ctx context.Context, record *ent.Book) *error2.AppError
	GetById(ctx context.Context, id string) (*ent.Book, *error2.AppError)
	DeleteById(ctx context.Context, id string) *error2.AppError
	Update(ctx context.Context, record *ent.Book) *error2.AppError
}
