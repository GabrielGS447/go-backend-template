package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Completed   bool               `json:"completed" bson:"completed"`
	Description string             `json:"description" bson:"description"`
	Date        string             `json:"date" bson:"date"`
}

type CreateTodoDTO struct {
	Title       string `json:"title" bson:"title" validate:"required,min=5,max=50"`
	Completed   *bool  `json:"completed" bson:"completed" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required,min=5,max=50"`
	Date        string `json:"date" bson:"date" validate:"required,date"`
}

type UpdateTodoDTO struct {
	Title       string `json:"title" bson:"title" validate:"required,min=5,max=50"`
	Completed   *bool  `json:"completed" bson:"completed" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required,min=5,max=50"`
	Date        string `json:"date" bson:"date" validate:"required,date"`
}
