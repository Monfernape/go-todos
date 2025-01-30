package main

import (
	"net/http"
	handlers "todos-api/handler"
)

var TodosPath = "/todos"

func main() {

	http.HandleFunc(TodosPath, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTodos(w, r)
		case http.MethodPost:
			handlers.CreateTodo(w, r)
		case http.MethodPut:
			handlers.UpdateTodo(w, r)
		case http.MethodDelete:
			handlers.DeleteTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
