package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"rest-api/internal/pkg/config"
)

var MongoConfig *mongo.Database

func ProviderNoSQL(environment config.Environment) (*mongo.Database, error) {
	uri := environment.Mongo.Connection
	client, err := mongo.Connect(options.Client().ApplyURI(uri).SetRetryWrites(false))
	if err != nil {
		panic("cannot connect to database")
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	MongoConfig = client.Database(environment.Mongo.Database)
	return MongoConfig, nil
}
