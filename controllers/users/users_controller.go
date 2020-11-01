package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorreis/bookstore-users-api/domain/users"
	"github.com/vitorreis/bookstore-users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// handle error
		fmt.Println("read all error")
		return
	}
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("s: bytes")
	fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &user); err != nil {
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
