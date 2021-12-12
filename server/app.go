package server

import (
	"github.com/gin-gonic/gin"
	"todolist/repositories/memory_impl"
	"todolist/server/handlers"
)

func RunApp() error {
	router := gin.Default()

	todoRepository := memory_impl.CreateMemoryTodoRepository()

	handlers.CreateTodoHandlers(router, todoRepository)

	return router.Run(":8000")
}
