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

func SeedOrder(c context.Context) error {
	var orderJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonFile, err := os.ReadFile("/home/mohammed/Desktop/selling_tmp/public/order.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(jsonFile, &orderJSON)

	if err != nil {
		fmt.Println(err.Error())
	}
	bulk := make([]*ent.OrderCreate, len(orderJSON))

	for i, od := range orderJSON {
		var order ent.Order
		err := mapstructure.Decode(od, &order)
		if err != nil {
			fmt.Println(err.Error())
		}
		bulk[i] = db.Client.Order.Create().
			SetUserID(order.UserID).
			SetProductID(order.ProductID).
			SetQuantity(order.Quantity).
			SetTotal(order.Total)
	}
	_, err = db.Client.Order.CreateBulk(bulk...).Save(c)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
