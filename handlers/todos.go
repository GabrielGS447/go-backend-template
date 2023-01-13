package handlers

import (
	"github.com/bmdavis419/go-backend-template/dtos"
	"github.com/bmdavis419/go-backend-template/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body dtos.CreateTodo true "Todo to create"
// @Produce json
// @Success 200 {object} dtos.CreateTodoRes
// @Router /todos [post]
func CreateTodo(c *fiber.Ctx) error {
	nTodo := new(dtos.CreateTodo)

	if err := c.BodyParser(nTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	insertedId, err := services.CreateTodo(c.Context(), nTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
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
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
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
	id := c.Locals("id").(primitive.ObjectID) // id is parsed by middleware

	todo, err := services.GetTodoById(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	if todo == nil {
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	}

	return c.Status(200).JSON(todo)
}

// @Summary Update a todo.
// @Description update a single todo.
// @Tags todos
// @Accept json
// @Param todo body dtos.UpdateTodo true "Todo update data"
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} dtos.UpdateTodoRes
// @Router /todos/:id [put]
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Locals("id").(primitive.ObjectID) // id is parsed by middleware

	uTodo := new(dtos.UpdateTodo)

	if err := c.BodyParser(uTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	modifiedCount, err := services.UpdateTodo(c.Context(), id, uTodo)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	if modifiedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "todo updated"})
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} dtos.DeleteTodoRes
// @Router /todos/:id [delete]
func DeleteTodo(c *fiber.Ctx) error {
	dbId := c.Locals("id").(primitive.ObjectID) // id is parsed by middleware

	deletedCount, err := services.DeleteTodo(c.Context(), dbId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	if deletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "todo deleted"})
}
