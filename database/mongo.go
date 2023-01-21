package database

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	todosCollection = "todos"
)

var mongoClient *mongo.Client
var dbName string

func getCollection(name string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(name)
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		return errors.New("you must set your 'DATABASE' environmental variable")
	} else {
		dbName = database
	}

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return nil
}

func CloseMongoDB() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}
