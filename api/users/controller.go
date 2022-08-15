package users

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get users with params(name, age)
func getUsers(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// post user with body request
func postUser(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/users")
	route.GET("/:name/:msg", getUsers)
	route.POST("/", postUser)
}
