package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todolist/repositories/interfaces"
)

type TodoHandlers struct {
	repo interfaces.TodoRepository
}

func CreateTodoHandlers(router *gin.Engine, repo interfaces.TodoRepository) TodoHandlers {
	handlers := TodoHandlers{
		repo: repo,
	}

	router.POST("/todos/", handlers.Add)
	router.GET("/todos/", handlers.GetAll)
	router.GET("/todos/:id", handlers.GetById)
	router.PUT("/todos/:id/complete", handlers.Complete)

	return handlers
}

func (handler *TodoHandlers) Add(ctx *gin.Context) {
	text := ctx.Query("text")
	todo, err := handler.repo.Add(text)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "An error occurred on the server",
		})
		return
	}

	ctx.JSONP(200, todo)
}

func (handler *TodoHandlers) GetAll(ctx *gin.Context) {
	todos, err := handler.repo.GetAll()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "An error occurred on the server",
		})
		return
	}

	ctx.JSONP(200, todos)
}

func (handler *TodoHandlers) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := handler.repo.GetById(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": fmt.Sprintf("Todo with id = %v not found", id),
		})
		return
	}

	ctx.JSONP(200, todo)
}

func (handler *TodoHandlers) Complete(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := handler.repo.Complete(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": fmt.Sprintf("Todo with id = %v not found", id),
		})
		return
	}

	ctx.JSONP(200, todo)
}
