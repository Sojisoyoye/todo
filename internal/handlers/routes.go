package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Sojisoyoye/todo/internal/service"
)

func RegisterRoutes(r *gin.Engine, todoService service.TodoServiceInterface) {
	todoHandler := NewTodoHandler(todoService)

	api := r.Group("/api/v1")

	{
		todos := api.Group("/todos")
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.GetAllTodos)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
	    }
	}
}