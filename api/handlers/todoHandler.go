package handlers

import (
	"EventManagement/api/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTodo Todo Add Handler
func AddTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewTodo()
		c.Bind(&requestBody)
		services.AddTodoService(requestBody)
		c.JSON(http.StatusOK, services.GetTodoList())
	}
}

// UpdateTodo Todo Update Handler
func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewTodo()
		c.Bind(&requestBody)
		services.UpdateTodoService(requestBody)
		c.JSON(http.StatusOK, services.GetTodoList())
	}
}

// DeleteTodo Todo Delete Handler
func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := services.NewTodo()
		c.Bind(&requestBody)
		services.DeleteTodoService(requestBody)
		c.JSON(http.StatusOK, services.GetTodoList())
	}
}

// ViewTodos Todo View Handler
func ViewTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetTodoList())
	}
}
