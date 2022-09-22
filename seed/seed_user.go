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

func SeedUser(ctx context.Context) error {
	var usersJSON []gin.H
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonFile, err := os.ReadFile("/home/mohammed/Desktop/selling_tmp/public/users.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully Opened MOCK_DATA.json")

	err = json.Unmarshal(jsonFile, &usersJSON)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, u := range usersJSON {
		var user ent.User
		err := mapstructure.Decode(u, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		db.Client.User.Create().
			SetName(user.Name).
			SetUsername(user.Username).
			SetPassword(user.Password).
			SetEmail(user.Email).
			SetAge(user.Age).
			SetPhone(user.Phone).
			SetNationalID(user.NationalID).
			SetLocalAddress(user.LocalAddress).
			SaveX(ctx)
	}
	fmt.Println("Successfully Created")

	return nil
}
