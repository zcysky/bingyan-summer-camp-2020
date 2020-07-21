package main

import (
	"github.com/gin-gonic/gin"
	"task1/controller"
)

func main() {
	r := gin.Default()

	r.POST("/registerUser", controller.Register)
	r.POST("/loginUser", controller.Login)
	r.POST("/listUser", controller.List)

	_ = r.Run(":8080")
}
