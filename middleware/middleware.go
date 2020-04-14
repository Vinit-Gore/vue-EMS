package middleware

import (
	"fmt"
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/api/helpers"
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/api/models"
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/api/routes"
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/api/services"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init -Init
func InitMiddleware(g *gin.Engine) {
	g.Use(cors.Default()) // CORS request
	fmt.Println("InitMiddleware called")

	o := g.Group("/o")
	o.Use(OpenRequestMiddleware())
	r := g.Group("/r")
	r.Use(RestrictedRequestMiddleware())
	// r.Use(jwt.Auth(models.JWTKey))
	c := r.Group("/c")
	c.Use(RoleBasedRequestMiddleware())
	routes.Init(o, r, c)
	routes.InitLoginRoute(o, r, c)

}

func OpenRequestMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("OpenRequestMiddleware called")
	}
}

// Need to check JWT token here
func RestrictedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		login, err := helpers.GetLoginFromToken(c)
		if err != nil {
			fmt.Println("Token not available", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		if strings.Trim(token, "") == "" {
			fmt.Println("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		//Todo: validate token
		user := models.User{}
		user.LoginID = login.LoginID
		user.Password = login.Password
		user, isValid, usererr := services.ValidatedUser(user)
		if usererr != nil || !isValid {
			fmt.Println("Failed to validate user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		}

		c.Next()
		fmt.Println("RestrictedRequestMiddleware called")
	}
}

// Need to check JWT token here with group validation
func RoleBasedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("RoleBasedRequestMiddleware called")
	}
}
