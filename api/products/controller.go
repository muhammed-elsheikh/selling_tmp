package products

import (
	"fmt"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"
	"selling_tmp/ent/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get Products
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

// Create Product
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

	c.JSON(http.StatusCreated, gin.H{"Prductes": product})
}

// Update Product
func updateProduct(c *gin.Context) {
	var inputs *ent.Product

	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := db.Client.Product.
		UpdateOneID(inputs.ID).
		SetName(inputs.Name).
		SetPrice(inputs.Price).
		SetSize(inputs.Size).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, product)
}

// Get Product
func getProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	product, err := db.Client.Product.Query().
		Where(product.ID(productID)).
		First(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{"product_details": product})
}

// Delete Product
func deleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	product := db.Client.Product.DeleteOneID(productID).Exec(c)

	c.JSON(http.StatusCreated, gin.H{"delete_product": product})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/products")
	route.GET("/all", getProducts)
	route.POST("", createProduct)
	route.PUT("/update", updateProduct)
	route.GET("/:id", getProduct)
	route.DELETE("/:id", deleteProduct)
}
