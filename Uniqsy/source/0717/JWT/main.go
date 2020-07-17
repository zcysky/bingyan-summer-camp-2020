package main

import (
	"JWT/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//获取token
	r.GET("/jwt/getToken", controller.GetToken)

	//解析token
	r.POST("/jwt/verifyToken", controller.VerifyToken)

	_ = r.Run(":8080")
}