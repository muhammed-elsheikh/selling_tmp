package orders

import (
	"fmt"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"

	"github.com/gin-gonic/gin"
)

// get orders with params(user, product)
func getOrders(c *gin.Context) {
	var outputs *ent.Order

	if err := c.BindJSON(&outputs); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	order, err := db.Client.Order.Query().All(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, order)
}

func createOrder(c *gin.Context) {
	var inputs *ent.Order

	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	order, err := db.Client.Order.Create().
		SetUserID(inputs.UserID).
		SetProductID(inputs.ProductID).
		SetQuantity(inputs.Quantity).
		SetTotal(inputs.Total).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusCreated, order)
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/orders")
	route.GET("/all", getOrders)
	route.POST("", createOrder)
}
