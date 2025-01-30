package handlers

import (
	"encoding/json"
	"net/http"
	"todos-api/logic"
	"todos-api/models"
	"todos-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	result, error := logic.GetTodos()
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result, error := logic.CreateTodo(todo)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := utils.GetIdFromPath(r.URL.Path)

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// typecase todoId to bson.ObjectId
	objID, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	todo.Id = objID

	result, error := logic.UpdateTodo(todo)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoId := utils.GetIdFromPath(r.URL.Path)
	deletedTodo, error := logic.DeleteTodo(todoId)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedTodo)
}
