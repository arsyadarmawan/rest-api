package ent

import "time"

type Book struct {
	ID          string    `bson:"_id" mapper:"ID"`
	Name        string    `bson:"name"`
	CreatedAt   time.Time `bson:",created_at"`
	UpdatedAt   time.Time `bson:",updated_at"`
	PublishedBy string    `bson:"publishedBy"`
	Author      string    `bson:"author"`
	ReleaseDate time.Time `bson:"releaseDate"`
	Description string    `bson:",notnull,omitempty"`
}
