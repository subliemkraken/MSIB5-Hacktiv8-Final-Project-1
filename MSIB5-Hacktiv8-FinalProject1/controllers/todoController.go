package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo represents the model for a todo
type Todo struct {
	ID   string `json:"id" example:"1"`
	Task string `json:"task" example:"tidur"`
}

var AllTodos = []Todo{}

// GetAllTodo godoc
// @Summary Get all todos
// @Description Get details of all todo list
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} controllers.Todo
// @Router /todo [get]
func GetAllTodo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success",
		"todo":   AllTodos,
	})
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a todo
// @Tags todo
// @Accept json
// @Produce json
// @Param controllers.Todo body controllers.Todo true "create a todo"
// @Success 200 {object} controllers.Todo
// @Router /todo [post]
func CreateTodo(ctx *gin.Context) {
	var newTodo Todo

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newTodo.ID = fmt.Sprintf("todo-%d", len(AllTodos)+1)
	AllTodos = append(AllTodos, newTodo)

	ctx.JSON(http.StatusCreated, gin.H{
		"todo": newTodo,
	})
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo with ID
// @Tags todo
// @Accept json
// @Produce json
// @Param todoID path string true "ID of todo to be updated"
// @Success 200 {object} controllers.Todo
// @Router /todo/{todoID} [put]
func UpdateTodo(ctx *gin.Context) {
	todoID := ctx.Param("todoID")
	condition := false
	var updatedTodo Todo

	if err := ctx.ShouldBindJSON(&updatedTodo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, todo := range AllTodos {
		if todoID == todo.ID {
			condition = true
			AllTodos[i] = updatedTodo
			AllTodos[i].ID = todoID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("todo with id %v not found.", todoID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("todo with id %v has been successfully updated", todoID),
	})
}

// GetTodoWithID godoc
// @Summary Get a todo
// @Description Get details of a todo
// @Tags todo
// @Accept json
// @Produce json
// @Param todoID path string true "ID of todo"
// @Success 200 {object} controllers.Todo
// @Router /todo/{todoID} [get]
func GetTodoWithID(ctx *gin.Context) {
	todoID := ctx.Param("todoID")
	condition := false
	var todoData Todo

	for i, todo := range AllTodos {
		if todoID == todo.ID {
			condition = true
			todoData = AllTodos[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("todo with id %v not found.", todoID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todo": todoData,
	})
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo with input todo's Id
// @Tags todo
// @Accept json
// @Produce json
// @Param todoID path string true "ID of todo to be deleted"
// @Success 204 "No Content"
// @Router /todo/{todoID} [delete]
func DeleteTodo(ctx *gin.Context) {
	todoID := ctx.Param("todoID")
	condition := false
	var todoIndex int

	for i, todo := range AllTodos {
		if todoID == todo.ID {
			condition = true
			todoIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("todo with id %v not found.", todoID),
		})
		return
	}

	copy(AllTodos[todoIndex:], AllTodos[todoIndex+1:])
	AllTodos[len(AllTodos)-1] = Todo{}
	AllTodos = AllTodos[:len(AllTodos)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("todo with id %v has been successfully deleted", todoID),
	})
}
