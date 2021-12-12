package models

import "fmt"

type Todo struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	IsCompleted bool   `json:"isCompleted"`
}

func (todo *Todo) String() string {
	return fmt.Sprintf(
		"Todo(Id=%v, Text=%v, IsCompleted=%v)",
		todo.Id, todo.Text, todo.IsCompleted,
	)
}

func CreateEmptyTodo() Todo {
	return Todo{Id: "", Text: "", IsCompleted: false}
}
