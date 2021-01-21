package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Login    string `json: "login`
	Password string `json: "password`
}

func main() {
	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/login", loginHandler)

	router.Run(":5000")
}

func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "ok",
	})
}

func loginHandler(c *gin.Context) {
	var form Login

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.user != "guilherme" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.Redirect(http.StatusOK, "/")
}
