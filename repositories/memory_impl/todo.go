package memory_impl

import (
	"github.com/google/uuid"
	"todolist/models"
	"todolist/repositories/errors"
)

type MemoryTodoRepository struct {
	todos []models.Todo
}

func CreateMemoryTodoRepository() *MemoryTodoRepository {
	return &MemoryTodoRepository{
		todos: make([]models.Todo, 0, 5),
	}
}

func (repo *MemoryTodoRepository) Add(text string) (models.Todo, error) {
	newTodo := models.Todo{Id: uuid.NewString(), Text: text, IsCompleted: false}
	repo.todos = append(repo.todos, newTodo)
	return newTodo, nil
}

func (repo *MemoryTodoRepository) GetById(id string) (models.Todo, error) {
	for _, todo := range repo.todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return models.CreateEmptyTodo(), &errors.TodoNotFound{Id: id}
}

func (repo *MemoryTodoRepository) GetAll() ([]models.Todo, error) {
	return repo.todos, nil
}

func (repo *MemoryTodoRepository) Complete(id string) (models.Todo, error) {
	for i := range repo.todos {
		if repo.todos[i].Id == id {
			repo.todos[i].IsCompleted = true
			return repo.todos[i], nil
		}
	}

	return models.CreateEmptyTodo(), &errors.TodoNotFound{Id: id}
}
