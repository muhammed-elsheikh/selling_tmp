package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"selling_tmp/ent"
	"selling_tmp/ent/migrate"
	"selling_tmp/ent/user"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func main() {
	client, err := ent.Open("mysql", "akkhor:Ma52569522??@tcp(127.0.0.1:3307)/selling_tmp?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to TabNine",
		})
	})

	router.GET("/index1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{
			"title": "Index1 website",
		})
	})
	router.GET("/index2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{
			"title": "Index2 website",
		})
	})
	router.GET("/index3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{
			"title": "Index3 website",
		})
	})
	router.GET("/index4", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{
			"title": "Index4 website",
		})
	})

	router.GET("/index5", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{
			"title": "Index5 website",
		})
	})

	router.Run(":3030")
}
