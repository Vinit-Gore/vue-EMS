package routes

import (
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/api/handlers"

	"github.com/gin-gonic/gin"
)

func Init(o, r, c *gin.RouterGroup) {

	r.POST("todos/add", handlers.AddTodo())
	r.POST("todos/update", handlers.UpdateTodo())
	o.GET("todos/view", handlers.ViewTodos())
	c.POST("todos/delete", handlers.DeleteTodo())
}
