package usecase

import "context"

type Book interface {
	Get(ctx context.Context) ([]BookResponse, error)
	GetById(ctx context.Context, id int) (BookResponse, error)
	Create(ctx context.Context, cmd BookRequest) error
	Delete(ctx context.Context, id string) error
}
