package protocol

import (
	"context"
	"todos-api/db"
	"todos-api/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DatabaseTable string

const (
	TodosTable DatabaseTable = "todos"
)

func List() ([]models.Todo, error) {
	mongoClient := db.GetMongoClient()
	cursor, err := mongoClient.Database("todos").Collection("todos").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var todos []models.Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func Create(input models.Todo) (interface{}, error)  {
	mongoClient := db.GetMongoClient()

	newTodo, err := mongoClient.Database(string(TodosTable)).Collection("todos").InsertOne(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

func Update(input models.Todo) (interface{}, error) {
	mongoClient := db.GetMongoClient()

	updatedTodo, err := mongoClient.Database(string(TodosTable)).Collection("todos").UpdateOne(context.Background(), bson.M{"_id": input.Id}, bson.M{"$set": input})
	if err != nil {
		return nil, err
	}
	return updatedTodo, nil
}

func Delete(todoId string) (interface{}, error) {
	mongoClient := db.GetMongoClient()

	deletedTodo, err := mongoClient.Database(string(TodosTable)).Collection("todos").DeleteOne(context.Background(), bson.M{"_id": todoId})
	if err != nil {
		return nil, err
	}
	return deletedTodo, nil
}
