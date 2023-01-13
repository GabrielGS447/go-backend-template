package services

/*
	Services are the middle layer between the handlers and the database.
	They are responsible for handling the business logic of the application.
	Not really applicable in this example app but it's good practice to have.
*/

import (
	"context"

	"github.com/bmdavis419/go-backend-template/database"
	"github.com/bmdavis419/go-backend-template/dtos"
	"github.com/bmdavis419/go-backend-template/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(ctx context.Context, data *dtos.CreateTodo) (primitive.ObjectID, error) {
	return database.CreateTodo(ctx, data)
}

func GetAllTodos(ctx context.Context) (*[]models.Todo, error) {
	return database.GetAllTodos(ctx)
}

func GetTodoById(ctx context.Context, id primitive.ObjectID) (*models.Todo, error) {
	return database.GetTodoById(ctx, id)
}

func UpdateTodo(ctx context.Context, id primitive.ObjectID, data *dtos.UpdateTodo) (int64, error) {
	return database.UpdateTodo(ctx, id, data)
}

func DeleteTodo(ctx context.Context, id primitive.ObjectID) (int64, error) {
	return database.DeleteTodo(ctx, id)
}
