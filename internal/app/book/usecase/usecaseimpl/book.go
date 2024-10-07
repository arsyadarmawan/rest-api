package usecaseimpl

import (
	"context"
	"rest-api/internal/app/book/repository"
	"rest-api/internal/app/book/usecase"
	"rest-api/internal/app/ent"
	"strconv"
	"time"
)

type BookOpts struct {
	Repository repository.BookRepository
}

type Book struct {
	Opts BookOpts
}

func NewBookImpl(opt BookOpts) *Book {
	return &Book{
		Opts: opt,
	}
}

func (b Book) Get(ctx context.Context) (records []usecase.BookResponse, err error) {
	bookRecords, err := b.Opts.Repository.Get(ctx)
	if err != nil {
		return []usecase.BookResponse{}, err
	}
	for _, record := range bookRecords {
		records = append(records, usecase.BookResponse{
			Name:        record.Name,
			Description: record.Description,
		})
	}
	return
}

func (b Book) GetById(ctx context.Context, id int) (usecase.BookResponse, error) {
	ent, err := b.Opts.Repository.GetById(ctx, strconv.Itoa(id))
	if err != nil {
		return usecase.BookResponse{}, err
	}

	return usecase.BookResponse{
		Name:        ent.Name,
		Description: ent.Description,
	}, nil
}

func (b Book) Create(ctx context.Context, cmd usecase.BookRequest) error {
	return b.Opts.Repository.Create(ctx, b.converCmdIntoEnt(cmd))
}

func (b Book) converCmdIntoEnt(cmd usecase.BookRequest) *ent.Book {
	return &ent.Book{
		Name:        cmd.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: cmd.Description,
	}
}
