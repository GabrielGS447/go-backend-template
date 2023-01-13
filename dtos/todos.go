package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTodo struct {
	Title       string `json:"title" bson:"title" validate:"required,min=5,max=50"`
	Completed   *bool  `json:"completed" bson:"completed" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required,min=5,max=50"`
	Date        string `json:"date" bson:"date" validate:"required,date"`
}

type CreateTodoRes struct {
	InsertedId primitive.ObjectID `json:"inserted_id" bson:"_id"`
}

type UpdateTodo struct {
	Title       string `json:"title" bson:"title" validate:"required,min=5,max=50"`
	Completed   *bool  `json:"completed" bson:"completed" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required,min=5,max=50"`
	Date        string `json:"date" bson:"date" validate:"required,date"`
}

type UpdateOrDeleteTodoRes struct {
	Message string `json:"message" bson:"message"`
}
