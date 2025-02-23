package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"todos-api/db"
	"todos-api/models"
	"todos-api/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DatabaseTable string

const (
	TodosTable DatabaseTable = "todos"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	mongoClient := db.GetMongoClient()
	cursor, err := mongoClient.Database("todos").Collection("todos").Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	defer cursor.Close(context.Background())

	var todos []models.Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to decode todos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	mongoClient := db.GetMongoClient()
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTodo, err := mongoClient.Database(string(TodosTable)).Collection("todos").InsertOne(context.Background(), todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	mongoClient := db.GetMongoClient()
	todoId := utils.GetIdFromPath(r.URL.Path)

	fmt.Sprintln("Updating todo with id: %s", todoId)

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	updatedTodo, err := mongoClient.Database(string(TodosTable)).Collection("todos").UpdateOne(context.Background(), bson.M{"_id": todoId}, bson.M{"$set": todo})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	mongoClient := db.GetMongoClient()

	todoId := utils.GetIdFromPath(r.URL.Path)
	fmt.Sprintln("Deleting todo with id: %s", todoId)

	deletedTodo, err := mongoClient.Database(string(TodosTable)).Collection("todos").DeleteOne(context.Background(), bson.M{"_id": todoId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedTodo)
}
