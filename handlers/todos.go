package handlers

import (
	"encoding/json"
	"fmt"

	_ "github.com/bmdavis419/go-backend-template/docs"
	"github.com/bmdavis419/go-backend-template/errs"
	"github.com/bmdavis419/go-backend-template/models"
	"github.com/bmdavis419/go-backend-template/services"
	"github.com/bmdavis419/go-backend-template/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body models.CreateTodoDTO true "Todo to create"
// @Produce json
// @Success 200 {object} CreateTodoRes
// @Router /todos [post]
func CreateTodo(c *fiber.Ctx) error {
	nTodo := new(models.CreateTodoDTO)

	if err := c.BodyParser(nTodo); err != nil {
		return handleTodoErrors(c, err)
	}

	if err := utils.ValidateInput(nTodo); err != nil {
		return handleTodoErrors(c, err)
	}

	insertedId, err := services.CreateTodo(c.Context(), nTodo)
	if err != nil {
		return handleTodoErrors(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"todo_id": insertedId})
}

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func GetAllTodos(c *fiber.Ctx) error {
	todos, err := services.GetAllTodos(c.Context())
	if err != nil {
		return handleTodoErrors(c, err)
	}

	return c.Status(200).JSON(todos)
}

// @Summary Get a single todo.
// @Description fetch a single todo.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos/:id [get]
func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := services.GetTodoById(c.Context(), id)
	if err != nil {
		return handleTodoErrors(c, err)
	}

	return c.Status(200).JSON(todo)
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
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	uTodo := new(models.UpdateTodoDTO)

	if err := c.BodyParser(uTodo); err != nil {
		return handleTodoErrors(c, err)
	}

	if err := utils.ValidateInput(uTodo); err != nil {
		return handleTodoErrors(c, err)
	}

	err := services.UpdateTodo(c.Context(), id, uTodo)
	if err != nil {
		return handleTodoErrors(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"message": "todo updated"})
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} UpdateOrDeleteTodoRes
// @Router /todos/:id [delete]
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.DeleteTodo(c.Context(), id)
	if err != nil {
		return handleTodoErrors(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"message": "todo deleted"})
}

func handleTodoErrors(c *fiber.Ctx, err error) error {
	if valErrs := utils.GetValidationErrors(err); valErrs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": valErrs})
	}

	switch err {
	case errs.ErrTodoNotFound:
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	case err.(*json.SyntaxError):
		return c.Status(422).JSON(fiber.Map{"error": "Invalid JSON syntax"})
	default:
		fmt.Println("error:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Something went wrong, please try again later."})
	}
}

type CreateTodoRes struct {
	InsertedId primitive.ObjectID `json:"inserted_id" bson:"_id"`
}

type UpdateOrDeleteTodoRes struct {
	Message string `json:"message" bson:"message"`
}
