package seed

import (
	"context"
	"fmt"
	"os"
	"selling_tmp/db"
	"selling_tmp/ent"

	"github.com/mitchellh/mapstructure"
)

func Do(ctx context.Context) error {
	var users []*ent.User

	jsonFile, err := os.Open("/home/mohammed/Desktop/selling_tmp/public/MOCK_DATA.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully Opened MOCK_DATA.json")

	err = mapstructure.Decode(jsonFile, &users)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, user := range users {
		db.Client.User.Create().
			SetName(user.Name).
			SetUsername(user.Username).
			SetPassword(user.Password).
			SetEmail(user.Email).
			SetAge(user.Age).
			SetNationalID(user.NationalID).
			SetLocalAddress(user.LocalAddress).
			SaveX(ctx)
	}
	fmt.Println("Successfully Created")

	return nil
}
