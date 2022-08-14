package users

import (
	"net/http"
	"selling_tmp/ent"

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

func createUser(c *gin.Context) {
	var user ent.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name != "manu" || user.Phone != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/users")
	route.GET("/:name/:msg", getUsers)
	route.POST("", createUser)
}
