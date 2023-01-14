package database

import (
	"context"
	"time"

	"github.com/bmdavis419/go-backend-template/errs"
	"github.com/bmdavis419/go-backend-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTodo(ctx context.Context, data *models.CreateTodoDTO) (string, error) {
	coll := getCollection(TodosCollection)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := coll.InsertOne(timeoutCtx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetAllTodos(ctx context.Context) (*[]models.Todo, error) {
	coll := getCollection(TodosCollection)

	filter := bson.M{}
	opts := options.Find().SetSkip(0).SetLimit(100)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := coll.Find(timeoutCtx, filter, opts)
	if err != nil {
		return nil, err
	}

	todos := make([]models.Todo, 0)

	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return &todos, nil
}

func GetTodoById(ctx context.Context, id string) (*models.Todo, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errs.ErrTodoNotFound
	}

	coll := getCollection(TodosCollection)
	filter := bson.M{"_id": objectId}
	todo := new(models.Todo)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := coll.FindOne(timeoutCtx, filter).Decode(todo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.ErrTodoNotFound
		} else {
			return nil, err
		}
	}

	return todo, nil
}

func UpdateTodo(ctx context.Context, id string, data *models.UpdateTodoDTO) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errs.ErrTodoNotFound
	}

	coll := getCollection(TodosCollection)
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": data}

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := coll.UpdateOne(timeoutCtx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errs.ErrTodoNotFound
	}

	return nil
}

func DeleteTodo(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errs.ErrTodoNotFound
	}

	coll := getCollection(TodosCollection)
	filter := bson.M{"_id": objectId}

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := coll.DeleteOne(timeoutCtx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errs.ErrTodoNotFound
	}

	return nil
}
