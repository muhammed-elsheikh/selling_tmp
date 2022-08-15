package main

import (
	"context"
	"fmt"
	"selling_tmp/api"
	"selling_tmp/db"
	"selling_tmp/ent"
	_ "selling_tmp/ent/runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func Do(ctx context.Context, client *ent.Client) error {
	_, err := client.User.
		Create().
		SetName("Mohammed").
		SetAge(20).
		SetEmail("mhdshaikh20403@gmail.com").
		SetPhone("011525695822").
		SetNationalID("225566998855").
		SetLocalAddress("Ahmed Elmenofy Street, Elsalam city, Cairo, Egypt").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = client.User.
		Create().
		SetName("Mr Yasser").
		SetAge(34).
		SetEmail("yseer.sobhy@gmail.com").
		SetPhone("01014896523").
		SetNationalID("33666985522").
		SetLocalAddress("Ahmed Elmenofy Street, Elsalam city, Cairo, Egypt").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}
	_, err = client.User.
		Create().
		SetName("Omar Abdo").
		SetAge(15).
		SetEmail("omar.abdo@gmail.com").
		SetPhone("01147958746").
		SetNationalID("369852147855").
		SetLocalAddress("Ahmed Elmenofy Street, Elsalam city, Cairo, Egypt").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}
	_, err = client.Card.
		Create().
		SetOwnerID(1).
		SetNumber("2563417847").
		SetExpiredTime(time.Now().Add(time.Minute)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating Card: %w", err)
	}
	_, err = client.Card.
		Create().
		SetUserID(2).
		SetNumber("4111225554").
		SetExpiredTime(time.Now().Add(time.Minute)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating Card: %w", err)
	}
	_, err = client.Card.
		Create().
		SetUserID(3).
		SetNumber("7778889655").
		SetExpiredTime(time.Now().Add(time.Minute)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating Card: %w", err)
	}

	_, err = client.Product.
		Create().
		SetName("Table").
		SetPrice(158.5).
		SetSerialNumber("12345").
		SetSize("Larg").
		SetCompanyName("Elsherbiny").
		SetShortDescription("Made of wood and u can use it to eat on it.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Product: %w", err)
	}

	_, err = client.Product.
		Create().
		SetName("Chair").
		SetPrice(99.75).
		SetSerialNumber("32165").
		SetSize("Medium").
		SetCompanyName("Elsherbiny").
		SetShortDescription("Made of wood and u can use it set on it.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Product: %w", err)
	}

	_, err = client.Product.
		Create().
		SetName("Office").
		SetPrice(122).
		SetSerialNumber("45876").
		SetSize("Small").
		SetCompanyName("Elsherbiny").
		SetShortDescription("Made of wood and u can use it study on it.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Product: %w", err)
	}

	_, err = client.Product.
		Create().
		SetName("Cup").
		SetPrice(12.5).
		SetSerialNumber("98754").
		SetSize("Medium").
		SetCompanyName("Elsherbiny").
		SetShortDescription("Made of wood and u can use it to drink on it.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Product: %w", err)
	}

	_, err = client.Order.
		Create().
		SetOwnerID(3).
		SetProductID(4).
		SetQuantity(5).
		SetTotal(452).
		SetOrderDate(time.Now()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Order: %w", err)
	}

	_, err = client.Order.
		Create().
		SetOwnerID(1).
		SetProductID(3).
		SetQuantity(4).
		SetTotal(788).
		SetOrderDate(time.Now()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Order: %w", err)
	}

	_, err = client.Order.
		Create().
		SetOwnerID(2).
		SetProductID(2).
		SetQuantity(3).
		SetTotal(669.5).
		SetOrderDate(time.Now()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Order: %w", err)
	}

	_, err = client.Order.
		Create().
		SetOwnerID(1).
		SetProductID(1).
		SetQuantity(12).
		SetTotal(1000.75).
		SetOrderDate(time.Now()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating Order: %w", err)
	}

	return nil
}

func main() {

	Do(context.Background(), db.Client)

	router := gin.Default()

	api.AddRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found!",
		})
	})

	router.Run(":3030")
}
