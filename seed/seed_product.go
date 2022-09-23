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

func SeedProduct(c context.Context) error {
	var productJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	jsonFile, err := os.ReadFile("/home/mohammed/Desktop/selling_tmp/public/products.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(jsonFile, &productJSON)

	if err != nil {
		fmt.Println(err.Error())
	}

	bulk := make([]*ent.ProductCreate, len(productJSON))
	for i, pd := range productJSON {
		var product ent.Product

		err := mapstructure.Decode(pd, &product)

		if err != nil {
			fmt.Println(err.Error())
		}

		bulk[i] = db.Client.Product.Create().
			SetName(product.Name).
			SetPrice(product.Price).
			SetSize(product.Size).
			SetSerialNumber(product.SerialNumber).
			SetCompanyName(product.CompanyName).
			SetShortDescription(product.ShortDescription)
	}

	_, err = db.Client.Product.CreateBulk(bulk...).Save(c)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully Created Products")
	return nil
}
