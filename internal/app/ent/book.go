package ent

import "time"

type Book struct {
	Name        string    `bson:"name"`
	CreatedAt   time.Time `bson:",created_at"`
	UpdatedAt   time.Time `bson:",updated_at"`
	Description string    `bson:",notnull,omitempty"`
}
