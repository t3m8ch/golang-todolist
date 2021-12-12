package interfaces

import (
	"todolist/models"
)

type TodoRepository interface {
	Add(text string) (models.Todo, error)
	GetById(id string) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Complete(id string) (models.Todo, error)
}
