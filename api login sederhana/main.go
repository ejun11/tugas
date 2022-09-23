package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/login", listLoginHandler)
	r.POST("/register", createRegisterHandler)

	r.Run()
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var logins = []Login{
	{Email: "bobi@gmail.com", Password: "ling"},
	{Email: "ucok@gmail.com", Password: "ucokaks"},
	{Email: "hari@gmail.com", Password: "hihihi"},
}

func listLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, logins)
}

func createRegisterHandler(c *gin.Context) {
	var login Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	logins = append(logins, login)

	c.JSON(http.StatusCreated, login)
}
