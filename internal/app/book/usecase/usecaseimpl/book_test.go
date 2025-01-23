package usecaseimpl_test

import (
	"context"
	"fmt"
	"github.com/arsyadarmawan/rest-api/internal/app/book/repository/repositorymock"
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase/usecaseimpl"
	"github.com/arsyadarmawan/rest-api/internal/app/ent"
	as "github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type mockBookRepo struct {
	repositorymock.MockBookRepository
}

func (b mockBookRepo) Get(ctx context.Context) ([]*ent.Book, error) {
	return []*ent.Book{
		{
			ID:          "1",
			Name:        "Test Book 1",
			Description: "Test Description 1",
		},
		{
			ID:          "2",
			Name:        "Test Book 2",
			Description: "Test Description 2",
		},
	}, nil
}

func (b mockBookRepo) GetById(ctx context.Context, id string) (*ent.Book, error) {
	return &ent.Book{
		ID:          "a3b7ff56-31c2-49b6-89fe-891867811f58",
		Name:        "Test Book 1",
		Description: "Test Description 1",
	}, nil
}

func (b mockBookRepo) DeleteById(ctx context.Context, id string) error {
	return nil
}

type mockBookRepoErr struct {
	repositorymock.MockBookRepository
}

func (b mockBookRepoErr) Get(ctx context.Context) ([]*ent.Book, error) {
	return nil, fmt.Errorf("simulated error")
}
func (b mockBookRepoErr) GetById(ctx context.Context, id string) (*ent.Book, error) {
	return nil, fmt.Errorf("simulated error")
}

func (b mockBookRepoErr) DeleteById(ctx context.Context, id string) error {
	return fmt.Errorf("simulated error")
}

func TestBook_Get(t *testing.T) {
	repo := &mockBookRepo{}
	book := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		Repository: repo,
	})

	resp, err := book.Get(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	want := []usecase.BookResponse{
		{
			Id:          "1",
			Name:        "Test Book 1",
			Description: "Test Description 1",
		},
		{
			Id:          "2",
			Name:        "Test Book 2",
			Description: "Test Description 2",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("Returned unexpected value: got %v want %v", resp, want)
	}
}

func TestBook_Get_Error(t *testing.T) {
	repo := &mockBookRepoErr{}
	book := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		Repository: repo,
	})

	_, err := book.Get(context.Background())
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestBook_GetByID(t *testing.T) {
	repo := &mockBookRepo{}
	book := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		Repository: repo,
	})

	resp, err := book.GetById(context.Background(), "a3b7ff56-31c2-49b6-89fe-891867811f58")
	if err != nil {
		t.Errorf("Expected an error, but got nil")
	}

	want := usecase.BookResponse{
		Id:          "a3b7ff56-31c2-49b6-89fe-891867811f58",
		Name:        "Test Book 1",
		Description: "Test Description 1",
	}

	assert := as.New(t)
	assert.Equal(resp.Id, want.Id)
	if !reflect.DeepEqual(resp, want) {
		t.Errorf("Returned unexpected value: got %v want %v", resp, want)
	}
}

func TestBook_GetByID_Err(t *testing.T) {
	repo := &mockBookRepoErr{}
	book := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		Repository: repo,
	})

	_, err := book.GetById(context.Background(), "a3b7ff56-31c2-49b6-89fe-891867811f58")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestBook_DeleteById(t *testing.T) {
	repo := &mockBookRepo{}
	book := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		Repository: repo,
	})
	err := book.Delete(context.Background(), "a3b7ff56-31c2-49b6-89fe-891867811f58")
	if err != nil {
		t.Errorf("Expected an error, but got nil")
	}
}
