package db

import (
	"context"
	"log"
	"selling_tmp/ent"
	"selling_tmp/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

var Client *ent.Client

func init() {
	var err error
	Client, err = ent.Open("mysql", "akkhor:Ma52569522??@tcp(localhost:3306)/selling_tmp")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer Client.Close()
	ctx := context.Background()
	err = Client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
