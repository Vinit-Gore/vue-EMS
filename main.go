package main

import (
	"fmt"
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin-Gonic Server")
	// append(todoList, {"Id":1,"Name":"Rohit"})
	startServer()
}

func startServer() {
	router := gin.Default()

	// gin.SetMode(gin.ReleaseMode)

	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	// md.ExposeHeaders = []string{"Authorization"}
	router.Use(cors.New(md))
	// g.Use(cors.Default())
	middleware.InitMiddleware(router)
	router.GET("/", checkStatus())

	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}
	s.ListenAndServe()
}

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
