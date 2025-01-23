package repository

import (
	"context"
	"github.com/arsyadarmawan/rest-api/internal/app/ent"
)

//go:generate mockgen -source=book.go -destination=repositorymock/book_mock.go -package=repositorymock
type BookRepository interface {
	Get(ctx context.Context) ([]*ent.Book, error)
	Create(ctx context.Context, record *ent.Book) error
	GetById(ctx context.Context, id string) (*ent.Book, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, record *ent.Book) error
}
