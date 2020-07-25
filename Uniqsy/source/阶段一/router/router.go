package router

import (
	"github.com/gin-gonic/gin"
	"task1/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/register", controller.Register)
	//r.POST("/api/login", controller.Login)
	//r.GET("/api/queryall", controller.QueryAll)
	//r.GET("/api/query", controller.Query)
	return r
}