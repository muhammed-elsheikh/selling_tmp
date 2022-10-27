package main

import (
	"selling_tmp/api"
	_ "selling_tmp/ent/runtime"

	"github.com/gin-gonic/gin"
)

func main() {

	// seed.SeedUser(context.Background())
	// seed.SeedCard(context.Background())
	// seed.SeedProduct(context.Background())
	// seed.SeedOrder(context.Background())

	router := gin.Default()

	api.AddRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found!",
		})
	})

	router.Run(":3031")
}
