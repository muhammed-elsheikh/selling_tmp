package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get products with params(product)
func getProducts(c *gin.Context) {
	product := c.Param("product")
	c.JSON(http.StatusOK, gin.H{
		"name": product,
	})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/products")
	route.GET("", getProducts)
}
