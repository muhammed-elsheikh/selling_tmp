package api

import (
	"selling_tmp/api/cards"
	"selling_tmp/api/orders"
	"selling_tmp/api/products"
	"selling_tmp/api/users"

	"github.com/gin-gonic/gin"
)

func AddRoutes(parentRoute *gin.Engine) {
	users.AddRoutes(parentRoute)
	products.AddRoutes(parentRoute)
	cards.AddRoutes(parentRoute)
	orders.AddRoutes(parentRoute)
}
