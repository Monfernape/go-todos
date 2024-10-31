package models

type Todo struct {
	Id        string `json:"_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
