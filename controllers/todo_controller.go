package controllers

import (
	"net/http"
	"strconv"

	"github.com/dalia3aly/go-sqlite-backend/models"
	"github.com/dalia3aly/go-sqlite-backend/services"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (c *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodoByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, err := c.todoService.GetTodoByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.todoService.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.todoService.UpdateTodo(uint(id), &todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.todoService.DeleteTodo(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
