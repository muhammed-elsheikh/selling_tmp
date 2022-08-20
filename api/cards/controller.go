package cards

import (
	"fmt"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"
	"selling_tmp/ent/card"

	"github.com/gin-gonic/gin"
)

// get crads with params(owner, num)
func getCards(c *gin.Context) {
	var outputs *ent.Card

	if err := c.BindJSON(&outputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cards, err := db.Client.Card.Query().All(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, cards)
}

func addCard(c *gin.Context) {

	var inputs *ent.Card

	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := db.Client.Card.Create().
		SetUserID(inputs.UserID).
		SetNumber(inputs.Number).
		SetExpiredTime(inputs.ExpiredTime).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusCreated, card)
}

func getCard(c *gin.Context) {
	var inputs *ent.Card

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := db.Client.Card.
		Query().
		Where(card.ID(inputs.ID)).
		First(c)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, card)
	}
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/cards")
	route.GET("/all", getCards)
	route.POST("", addCard)
	// route.POST("", getCard)
}
