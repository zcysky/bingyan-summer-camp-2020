package main

import (
	"JWT/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/jwt/getToken", controller.GetToken)
	r.POST("/jwt/verifyToken", controller.VerifyToken)

	_ = r.Run(":8080")
}