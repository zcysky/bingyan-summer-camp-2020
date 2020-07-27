package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
)

// InitRouter creates a gin engine, sets up its route rules
// and then return the pointer of the engine
func InitRouter() *gin.Engine {
	router := gin.Default()
	userGroup := router.Group("")
	{
		userGroup.POST("/user", controller.Signup)
		userGroup.POST("/user/login", controller.Login)
		//get self info:
		userGroup.GET("/me", controller.LoginVerification, controller.GetSelfInfo)
		//update self info:
		userGroup.POST("/me", controller.LoginVerification, controller.Update)
		//get other user's info:
		userGroup.GET("/user/:id", controller.LoginVerification, controller.GetPublicInfo)
	}

	commodityGroup := router.Group("")
	{
		//get public commodity:
		commodityGroup.GET("/commodities", controller.GetCommodities)
		//get hot commodity keywords:
		commodityGroup.GET("/commodities/hot", controller.GetHots)
		//get self published commodity:
		commodityGroup.GET("/me/commodities", controller.LoginVerification, controller.QuerySelfCommodities)
		//to publish commodity:
		commodityGroup.POST("/commodities")
		//delete commodity:
		commodityGroup.DELETE("/commodity/:id", controller.LoginVerification)

	}

	collectionGroup := router.Group("")
	{
		collectionGroup.GET("/me/collections", controller.LoginVerification)    //get self collections
		collectionGroup.POST("/me/collections", controller.LoginVerification)   //add collection
		collectionGroup.DELETE("/me/collections", controller.LoginVerification) //delete collection
	}

	return router
}
