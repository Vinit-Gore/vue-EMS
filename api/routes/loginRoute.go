package routes

import (
	"EventManagement/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitLoginRoute(o, r, c *gin.RouterGroup) {
	o.POST("/login/registeruser", handlers.RegisterUser())
	o.POST("/login/loginuser", handlers.LoginUser())
}
