package products

import (
	"fmt"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"

	"github.com/gin-gonic/gin"
)

// get products
func getProducts(c *gin.Context) {
	var outputs *ent.Product

	if err := c.BindJSON(&outputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := db.Client.Product.Query().All(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var inputs *ent.Product
	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	product, err := db.Client.Product.Create().
		SetName(inputs.Name).
		SetPrice(inputs.Price).
		SetSize(inputs.Size).
		SetSerialNumber(inputs.SerialNumber).
		SetShortDescription(inputs.ShortDescription).
		SetCompanyName(inputs.CompanyName).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusCreated, product)
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/products")
	route.GET("/all", getProducts)
	route.POST("", createProduct)
}
