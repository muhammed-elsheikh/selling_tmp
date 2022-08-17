package users

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"selling_tmp/db"
	"selling_tmp/ent"
	"selling_tmp/ent/user"

	"github.com/gin-gonic/gin"
)

// get users with params(name, age)
func getUsers(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// post user with body request
func postUser(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}

func register(c *gin.Context) {
	var inputs *ent.User

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.Client.User.Create().
		SetName(inputs.Name).
		SetPassword(inputs.Password).
		SetAge(inputs.Age).
		SetEmail(inputs.Email).
		Save(c)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, user)
}

func login(c *gin.Context) {
	var inputs *ent.User

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.Client.User.
		Query().
		Where(user.Email(inputs.Email)).
		First(c)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func AddRoutes(parentRoute *gin.Engine) {
	route := parentRoute.Group("/users")
	route.GET("/:name/:msg", getUsers)
	route.POST("/", postUser)
	route.POST("/register", register)
	route.POST("/login", login)
}
