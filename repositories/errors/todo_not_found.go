package errors

import "fmt"

type TodoNotFound struct {
	Id string
}

func (error *TodoNotFound) Error() string {
	return fmt.Sprintf("Todo with id - %v not found", error.Id)
}
