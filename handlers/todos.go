package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/gabrielgs449/go-backend-template/docs"
	"github.com/gabrielgs449/go-backend-template/errs"
	"github.com/gabrielgs449/go-backend-template/models"
	"github.com/gabrielgs449/go-backend-template/services"
	"github.com/gabrielgs449/go-backend-template/utils"
	"github.com/labstack/echo/v4"
)

type TodosHandlerInterface interface {
	CreateTodo(c echo.Context) error
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
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
func (h *todosHandler) CreateTodo(c echo.Context) error {
	nTodo := new(models.CreateTodoDTO)

	if err := c.Bind(nTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	if err := utils.ValidateInput(nTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	insertedId, err := h.todosService.CreateTodo(c.Request().Context(), nTodo)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.JSON(http.StatusOK, CreateTodoRes{TodoId: insertedId})
}

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func (h *todosHandler) GetAllTodos(c echo.Context) error {
	todos, err := h.todosService.GetAllTodos(c.Request().Context())
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.JSON(http.StatusOK, todos)
}

// @Summary Get a single todo.
// @Description fetch a single todo.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos/:id [get]
func (h *todosHandler) GetTodoById(c echo.Context) error {
	id := c.Param("id")

	todo, err := h.todosService.GetTodoById(c.Request().Context(), id)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.JSON(http.StatusOK, todo)
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
func (h *todosHandler) UpdateTodo(c echo.Context) error {
	id := c.Param("id")

	uTodo := new(models.UpdateTodoDTO)

	if err := c.Bind(uTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	if err := utils.ValidateInput(uTodo); err != nil {
		return handleTodosErrors(c, err)
	}

	err := h.todosService.UpdateTodo(c.Request().Context(), id, uTodo)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.JSON(http.StatusOK, UpdateOrDeleteTodoRes{Message: fmt.Sprintf("todo %s updated", id)})
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} UpdateOrDeleteTodoRes
// @Router /todos/:id [delete]
func (h *todosHandler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	err := h.todosService.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return handleTodosErrors(c, err)
	}

	return c.JSON(http.StatusOK, UpdateOrDeleteTodoRes{Message: fmt.Sprintf("todo %s deleted", id)})
}

func handleTodosErrors(c echo.Context, err error) error {
	if valErrs := utils.GetValidationErrors(err); valErrs != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"errors": valErrs})
	}

	switch err {
	case errs.ErrTodoNotFound:
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	case err.(*json.SyntaxError):
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "Invalid JSON syntax"})
	default:
		fmt.Println("error:", err)
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": "Something went wrong, please try again later."},
		)
	}
}

type CreateTodoRes struct {
	TodoId string `json:"todo_id"`
}

type UpdateOrDeleteTodoRes struct {
	Message string `json:"message"`
}
