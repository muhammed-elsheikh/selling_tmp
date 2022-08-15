package cards

import (
	"net/http"
	"selling_tmp/ent"

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

func addCard(c *gin.Context) {

	var card ent.Card

	if err := c.BindJSON(&card); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, card)
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/cards")
	route.GET("", getCards)
	route.POST("", addCard)
}
