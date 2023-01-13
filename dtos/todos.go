package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateTodo struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
}

type CreateTodoRes struct {
	InsertedId primitive.ObjectID `json:"inserted_id" bson:"_id"`
}

type UpdateTodo struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
}

type UpdateTodoRes struct {
	UpdatedCount int64 `json:"updated_count" bson:"updated_count"`
}

type DeleteTodoRes struct {
	DeletedCount int64 `json:"deleted_count" bson:"deleted_count"`
}
