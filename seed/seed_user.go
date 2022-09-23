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
	bulk := make([]*ent.UserCreate, len(usersJSON))
	for i, u := range usersJSON {
		var user ent.User
		err := mapstructure.Decode(u, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		bulk[i] = db.Client.User.Create().
			SetName(user.Name).
			SetUsername(user.Username).
			SetPassword(user.Password).
			SetEmail(user.Email).
			SetAge(user.Age).
			SetPhone(user.Phone).
			SetNationalID(user.NationalID).
			SetLocalAddress(user.LocalAddress)
	}
	_, err = db.Client.User.CreateBulk(bulk...).Save(ctx)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully Created")

	return nil
}
