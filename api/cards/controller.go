package cards

import (
	"fmt"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"
	"selling_tmp/ent/card"
	"strconv"

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
	cardID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	card, err := db.Client.Card.Query().Where(card.ID(cardID)).First(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{"cardDetails": card})
}

func updateCard(c *gin.Context) {
	var inputs *ent.Card
	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	card, err := db.Client.Card.UpdateOneID(inputs.ID).
		SetUserID(inputs.UserID).
		SetNumber(inputs.Number).
		Save(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, card)
}

func deleteCard(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	card := db.Client.Card.DeleteOneID(cardID).Exec(c)

	c.JSON(http.StatusCreated, gin.H{"deleteCard": card})
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/cards")
	route.GET("/all", getCards)
	route.POST("", addCard)
	route.PUT("/update", updateCard)
	route.GET("/:id", getCard)
	route.DELETE("/:id", deleteCard)
}
