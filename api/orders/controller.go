package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get orders with params(user, product)
func getorders(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/products")
	route.GET("", getorders)
}
