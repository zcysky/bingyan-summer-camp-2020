package main

import (
	"github.com/gin-gonic/gin"
)

func sayhello(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "hello, king crimson",
	})
}


func main() {
	r := gin.Default()

	r.GET("/hello", sayhello)

	r.GET("/test", func(c *gin.Context){
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/test", func(c *gin.Context){
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/test", func(c *gin.Context){
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/test", func(c *gin.Context){
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	r.Run(":9090")
}