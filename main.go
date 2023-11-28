package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represent a user in the system
type User struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

// Users slice to add new user
var users []User

func main() {

	router := InitRouter()
	router.Run(":3000")

}

// Initializes a gin router and return it
func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.POST("user", CreateUser)
		v1.GET("user/:id", GetUserByID)
		v1.GET("users", GetUsers)
		v1.DELETE("user/:id", DeleteUserByID)
	}

	return r
}

// Create a new user in the system
func CreateUser(c *gin.Context) {

	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}

	users = append(users, newUser)
	fmt.Println(users)
	c.IndentedJSON(http.StatusCreated, newUser)

}

// Fetch all the user present in the system
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// Fetch user using user id
func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not present"})
}

// Delete user from the sytem based on user id
func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user with ID %s deleted", id)})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user with ID %s deleted", id)})
}
