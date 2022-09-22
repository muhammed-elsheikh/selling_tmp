package seed

import (
	"context"
	"fmt"
	"os"
	"selling_tmp/db"
	"selling_tmp/ent"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

func SeedCard(c context.Context) error {
	var cardJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonFile, err := os.ReadFile("/home/mohammed/Desktop/selling_tmp/public/cards.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(jsonFile, &cardJSON)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, cd := range cardJSON {
		var card ent.Card
		err := mapstructure.Decode(cd, &card)

		if err != nil {
			fmt.Println(err.Error())
		}

		db.Client.Card.Create().
			SetUserID(card.UserID).
			SetNumber(card.Number).
			SaveX(c)
	}
	fmt.Println("Successfully Created Cards")

	return nil
}
