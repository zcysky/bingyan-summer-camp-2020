package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		au := c.Request.Header.Get("Authorization")
		if au == "" {
			c.AbortWithStatus(401)
		} else {
			claim := ParseToken(au[7:])
			c.JSON(200, gin.H{
				"status":  "200 OK",
				"message": "Hello, " + claim["username"].(string),
			})
		}
	})
	r.Run(":3939")

	// jwt testing
	/*
		user := config.User{Username: "TobyZT", Password: "19260817"}
		str := GenerateToken(user)
		claim := ParseToken(str)
		fmt.Println("Hello, " + claim["username"].(string))
	*/
}
