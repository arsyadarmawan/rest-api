package repository

import (
	"context"
	"rest-api/internal/app/ent"
)

type BookRepository interface {
	Get(ctx context.Context) ([]*ent.Book, error)
	Create(ctx context.Context, record *ent.Book) error
	GetById(ctx context.Context, id string) (*ent.Book, error)
}
