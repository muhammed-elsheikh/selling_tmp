package users

import (
	"net/http"
	"selling_tmp/ent"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	name := c.Param("Moahammed")
	msg := name + "welcome"
	c.IndentedJSON(http.StatusOK, msg)
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
	route.GET("/", getUsers)
	route.POST("", createUser)
}
