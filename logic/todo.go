package logic

import (
	"todos-api/models"
	"todos-api/protocol"
)


func GetTodos() ([]models.Todo, error) {
	return protocol.List()
}

func CreateTodo(input models.Todo) (interface{}, error) {
	return protocol.Create(input)
}

func UpdateTodo(input models.Todo) (interface{}, error) {
	return protocol.Update(input)
}

func DeleteTodo(todoId string) (interface{}, error) {
	return protocol.Delete(todoId)
}
