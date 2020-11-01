package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorreis/bookstore-users-api/domain/users"
	"github.com/vitorreis/bookstore-users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// handle error
		fmt.Println("unmarshal error")
		return
	}
	fmt.Println(user)
	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		// handle error
		fmt.Println("save error")
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "implement me!",
	})
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "implement me!",
	})
}
