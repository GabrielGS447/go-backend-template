package handlers

import (
	"encoding/json"
	"fmt"

	_ "github.com/gabrielgs449/go-backend-template/docs"
	"github.com/gabrielgs449/go-backend-template/errs"
	"github.com/gabrielgs449/go-backend-template/models"
	"github.com/gabrielgs449/go-backend-template/services"
	"github.com/gabrielgs449/go-backend-template/utils"
	"github.com/gofiber/fiber/v2"
)

type TodosHandlerInterface interface {
	CreateTodo(c *fiber.Ctx) error
	GetAllTodos(c *fiber.Ctx) error
	GetTodoById(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
}

type todosHandler struct {
	todosService services.TodosServiceInterface
}

func NewTodoHandler(s services.TodosServiceInterface) TodosHandlerInterface {
	return &todosHandler{
		s,
	}
}

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body models.CreateTodoDTO true "Todo to create"
// @Produce json
// @Success 200 {object} CreateTodoRes
// @Router /todos [post]
func (h *todosHandler) CreateTodo(c *fiber.Ctx) error {
	nTodo := new(models.CreateTodoDTO)

	if err := c.BodyParser(nTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	if err := utils.ValidateInput(nTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	insertedId, err := h.todosService.CreateTodo(c.Context(), nTodo)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(CreateTodoRes{TodoId: insertedId})
}

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func (h *todosHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := h.todosService.GetAllTodos(c.Context())
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

// @Summary Get a single todo.
// @Description fetch a single todo.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos/:id [get]
func (h *todosHandler) GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := h.todosService.GetTodoById(c.Context(), id)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

// @Summary Update a todo.
// @Description update a single todo.
// @Tags todos
// @Accept json
// @Param todo body models.UpdateTodoDTO true "Todo update data"
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} UpdateOrDeleteTodoRes
// @Router /todos/:id [put]
func (h *todosHandler) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	uTodo := new(models.UpdateTodoDTO)

	if err := c.BodyParser(uTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	if err := utils.ValidateInput(uTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	err := h.todosService.UpdateTodo(c.Context(), id, uTodo)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(
		UpdateOrDeleteTodoRes{Message: fmt.Sprintf("todo with id %s updated", id)},
	)
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} UpdateOrDeleteTodoRes
// @Router /todos/:id [delete]
func (h *todosHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.todosService.DeleteTodo(c.Context(), id)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(
		UpdateOrDeleteTodoRes{Message: fmt.Sprintf("todo with id %s deleted", id)},
	)
}

func handleTodosErrors(c *fiber.Ctx, err error) error {
	if valErrs := utils.GetValidationErrors(err); valErrs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": valErrs})
	}

	switch err {
	case errs.ErrTodoNotFound:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	case err.(*json.SyntaxError):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Invalid JSON syntax"})
	default:
		fmt.Println("error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": "Something went wrong, please try again later."},
		)
	}
}

type CreateTodoRes struct {
	TodoId string `json:"todo_id"`
}

type UpdateOrDeleteTodoRes struct {
	Message string `json:"message"`
}
