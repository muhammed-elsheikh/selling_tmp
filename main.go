package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"selling_tmp/ent"
	"selling_tmp/ent/user"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
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
