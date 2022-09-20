package main

import (
	"context"
	"selling_tmp/api"
	_ "selling_tmp/ent/runtime"
	"selling_tmp/seed"

	"github.com/gin-gonic/gin"
)

func main() {

	seed.Do(context.Background())

	router := gin.Default()

	api.AddRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found!",
		})
	})

	router.Run(":3030")
}
