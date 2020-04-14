package handlers

import (
	"EventManagement/api/helpers"
	"EventManagement/api/models"
	"EventManagement/api/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		c.Bind(&user)
		services.AddUsers(user)
		fmt.Println(user)
		c.JSON(http.StatusOK, services.GetUserList())
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		user := models.User{}
		c.Bind(&user)
		user, isValidUser, err := services.ValidatedUser(user)
		if err != nil {
			fmt.Println("Invalid user!", err)
			// return
		}
		fmt.Println("isValidUser", isValidUser)
		if isValidUser {
			fmt.Println("LoginUser")
			token, err := helpers.GenerateToken(user.LoginID, user.Password, user.Group, 24*time.Hour)
			if err != nil {
				fmt.Print("error while generating token:", err)
				// return err
			}
			fmt.Println("Token:", token)

			c.Header("Authorization", token)
		}
		c.JSON(http.StatusOK, services.GetUserList())
	}
}
