package database

import (
	"context"

	"github.com/bmdavis419/go-backend-template/dtos"
	"github.com/bmdavis419/go-backend-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTodo(ctx context.Context, data *dtos.CreateTodo) (primitive.ObjectID, error) {
	coll := getCollection("todos")

	res, err := coll.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func GetAllTodos(ctx context.Context) (*[]models.Todo, error) {
	coll := getCollection("todos")

	filter := bson.M{}
	opts := options.Find().SetSkip(0).SetLimit(100)

	todos := make([]models.Todo, 0)

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return &todos, nil
}

func GetTodoById(ctx context.Context, id primitive.ObjectID) (*models.Todo, error) {
	coll := getCollection("todos")
	filter := bson.M{"_id": id}
	todo := new(models.Todo)

	if err := coll.FindOne(ctx, filter).Decode(todo); err != nil {
		// check if the error is a mongo.ErrNoDocuments
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return todo, nil
}

func UpdateTodo(ctx context.Context, id primitive.ObjectID, data *dtos.UpdateTodo) (int64, error) {
	coll := getCollection("todos")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}

	res, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return res.MatchedCount, nil
}

func DeleteTodo(ctx context.Context, id primitive.ObjectID) (int64, error) {
	coll := getCollection("todos")
	filter := bson.M{"_id": id}
	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return res.DeletedCount, nil
}
