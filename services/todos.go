package services

/*
	Services are the middle layer between the handlers and the database.
	They are responsible for handling the business logic of the application.
	Not really applicable in this example app but it's good practice to have.
*/

import (
	"context"

	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/gabrielgs449/go-backend-template/models"
)

type TodosServiceInterface interface {
	CreateTodo(ctx context.Context, data *models.CreateTodoDTO) (string, error)
	GetAllTodos(ctx context.Context) (*[]models.Todo, error)
	GetTodoById(ctx context.Context, id string) (*models.Todo, error)
	UpdateTodo(ctx context.Context, id string, data *models.UpdateTodoDTO) error
	DeleteTodo(ctx context.Context, id string) error
}

type todosService struct {
	todosRepository database.TodosRepositoryInterface
}

func NewTodosService(r database.TodosRepositoryInterface) TodosServiceInterface {
	return &todosService{
		r,
	}
}

func (s *todosService) CreateTodo(ctx context.Context, data *models.CreateTodoDTO) (string, error) {
	return s.todosRepository.CreateTodo(ctx, data)
}

func (s *todosService) GetAllTodos(ctx context.Context) (*[]models.Todo, error) {
	return s.todosRepository.GetAllTodos(ctx)
}

func (s *todosService) GetTodoById(ctx context.Context, id string) (*models.Todo, error) {
	return s.todosRepository.GetTodoById(ctx, id)
}

func (s *todosService) UpdateTodo(ctx context.Context, id string, data *models.UpdateTodoDTO) error {
	return s.todosRepository.UpdateTodo(ctx, id, data)
}

func (s *todosService) DeleteTodo(ctx context.Context, id string) error {
	return s.todosRepository.DeleteTodo(ctx, id)
}
