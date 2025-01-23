package usecaseimpl

import (
	"context"
	"fmt"
	"github.com/arsyadarmawan/asynq-distributed-task/enqueue"
	"github.com/arsyadarmawan/rest-api/internal/app/book/repository"
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
	"github.com/arsyadarmawan/rest-api/internal/app/ent"
	"github.com/arsyadarmawan/rest-api/internal/pkg/commonval"
	"github.com/google/uuid"
	"time"
)

type BookOpts struct {
	Repository    repository.BookRepository
	AsynqEnqueuer enqueue.AsynqEnqueuer
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
			Id:          record.ID,
			Name:        record.Name,
			Description: record.Description,
		})
	}
	return
}

func (b Book) GetById(ctx context.Context, id string) (usecase.BookResponse, error) {
	ent, err := b.Opts.Repository.GetById(ctx, id)
	if err != nil {
		return usecase.BookResponse{}, err
	}

	return usecase.BookResponse{
		Id:          ent.ID,
		Name:        ent.Name,
		Description: ent.Description,
	}, nil
}

func (b Book) Create(ctx context.Context, cmd usecase.BookRequest) error {
	ent := b.convertCmdIntoEnt(cmd)
	if err := b.Opts.AsynqEnqueuer.Enqueue(ctx, commonval.BookWorkerAsynq, ent.ID); err != nil {
		return err
	}
	return b.Opts.Repository.Create(ctx, ent)
}

func (b Book) Delete(ctx context.Context, id string) error {
	return b.Opts.Repository.DeleteById(ctx, id)
}

func (b Book) convertCmdIntoEnt(cmd usecase.BookRequest) *ent.Book {
	dateStr := "2025-01-14"
	releaseDateFormat, _ := time.Parse(cmd.ReleaseDate, dateStr)

	return &ent.Book{
		ID:          uuid.New().String(),
		Name:        cmd.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		PublishedBy: cmd.PublishedBy,
		Author:      cmd.Author,
		ReleaseDate: releaseDateFormat,
		Description: cmd.Description,
	}
}

func (b Book) Update(ctx context.Context, id string) error {
	bookRecord, err := b.Opts.Repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	bookRecord.Name = fmt.Sprintf("Updated %s", bookRecord.Name)
	bookRecord.UpdatedAt = time.Now()
	return b.Opts.Repository.Update(ctx, bookRecord)
}
