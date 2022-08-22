package orders

import (
	"fmt"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"
	"selling_tmp/ent/order"
	"strconv"

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

func getOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	order, err := db.Client.Order.Query().Where(order.ID(orderID)).First(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_details": order})
}

func updateOrder(c *gin.Context) {
	var inputs *ent.Order

	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	order, err := db.Client.Order.UpdateOneID(inputs.ID).
		SetUserID(inputs.UserID).
		SetProductID(inputs.ProductID).
		SetQuantity(inputs.Quantity).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, order)
}

func deleteOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	order := db.Client.Order.DeleteOneID(orderID)

	c.JSON(http.StatusOK, gin.H{"delete_order": order})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/orders")
	route.GET("/all", getOrders)
	route.POST("", createOrder)
	route.GET("/:id", getOrder)
	route.PUT("/update", updateOrder)
	route.DELETE("/:id", deleteOrder)
}
