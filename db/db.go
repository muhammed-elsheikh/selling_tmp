package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"selling_tmp/ent"
	"selling_tmp/ent/migrate"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var Client *ent.Client

func init() {
	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbName)
	Client, err = ent.Open(Dbdriver, DBURL)
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
