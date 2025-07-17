package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukegrady1/to-do-list-api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	todos := r.Group("/todos")
	{
		todos.GET("", handlers.GetTodos)
		todos.POST("", handlers.CreateTodo)
		todos.GET("/:id", handlers.GetTodoByID)
		todos.PUT("/:id", handlers.UpdateTodo)
		todos.DELETE("/:id", handlers.DeleteTodo)
	}

	return r
}
