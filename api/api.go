package api

import (
	"selling_tmp/api/users"

	"github.com/gin-gonic/gin"
)

func AddRoutes(parentRoute *gin.Engine) {
	users.AddRoutes(parentRoute)
}
