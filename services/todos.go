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
)

func CreateTodo(ctx context.Context, data *dtos.CreateTodo) (string, error) {
	return database.CreateTodo(ctx, data)
}

func GetAllTodos(ctx context.Context) (*[]models.Todo, error) {
	return database.GetAllTodos(ctx)
}

func GetTodoById(ctx context.Context, id string) (*models.Todo, error) {
	return database.GetTodoById(ctx, id)
}

func UpdateTodo(ctx context.Context, id string, data *dtos.UpdateTodo) error {
	return database.UpdateTodo(ctx, id, data)
}

func DeleteTodo(ctx context.Context, id string) error {
	return database.DeleteTodo(ctx, id)
}
