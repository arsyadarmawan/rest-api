package repositoryimpl

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"rest-api/internal/app/ent"
)

type BookRepositoryOpts struct {
	DB *mongo.Database
}

type BookRepository struct {
	Opts BookRepositoryOpts
}

func NewBookRepository(opts BookRepositoryOpts) *BookRepository {
	return &BookRepository{
		Opts: opts,
	}
}

func (b BookRepository) collectionName() *mongo.Collection {
	return b.Opts.DB.Collection("books")
}

func (b BookRepository) Get(ctx context.Context) (books []*ent.Book, err error) {
	cursor, err := b.collectionName().Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book ent.Book
		if err := cursor.Decode(&book); err != nil {
			log.Fatal(err)
		}
		books = append(books, &book)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func (b BookRepository) GetById(ctx context.Context, id string) (record *ent.Book, err error) {
	errCode := b.collectionName().FindOne(ctx, bson.D{}).Decode(record)
	if errCode != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = errors.New("document not found")
			return
		}
	}
	return
}

func (b BookRepository) Create(ctx context.Context, record *ent.Book) error {
	collection := b.collectionName()
	_, err := collection.InsertOne(ctx, record)
	if err != nil {
		return errors.New("error")
	}
	return nil
}

func (b BookRepository) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := b.collectionName().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
