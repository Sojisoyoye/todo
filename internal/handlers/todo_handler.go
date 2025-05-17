package handlers

import (
	"errors"
	"net/http"

	"github.com/Sojisoyoye/todo/internal/models"
	"github.com/Sojisoyoye/todo/internal/service"
	"github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type TodoHandler struct {
	TodoService service.TodoServiceInterface
}

func NewTodoHandler(todoService service.TodoServiceInterface) *TodoHandler {
	return &TodoHandler{
		TodoService: todoService,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.TodoService.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.TodoService.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    todo, err := h.TodoService.GetTodoByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    todo.ID = id

    err = h.TodoService.UpdateTodo(&todo)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        }
        return
    }

    c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    if err := h.TodoService.DeleteTodo(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}




