package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get crads with params(owner, num)
func getCards(c *gin.Context) {
	owner := c.Param("owner")
	num := c.Param("num")
	c.JSON(http.StatusOK, gin.H{
		"name":     owner,
		"card num": num,
	})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/products")
	route.GET("", getCards)
}
