package usecase

import "context"

//go:generate mockgen -source=book.go -destination=usecasemock/book_mock.go -package=usecasemock
type Book interface {
	Get(ctx context.Context) ([]BookResponse, error)
	GetById(ctx context.Context, id string) (BookResponse, error)
	Create(ctx context.Context, cmd BookRequest) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) error
}
