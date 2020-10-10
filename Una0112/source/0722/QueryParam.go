package main

//Querystring parameters

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Mr.") //如果没有值，还可以给一个默认值
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s ", firstname, lastname)
	})
	r.Run(":8080")
}