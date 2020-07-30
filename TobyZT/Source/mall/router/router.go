package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
)

// InitRouter creates a gin engine, sets up its route rules
// and then return the pointer of the engine
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 4 << 20 // limit: 4MB
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
		//get particular commodity:
		commodityGroup.GET("/commodity/:id", controller.TryToGetUser, controller.GetCommodityByID)
		//get self published commodity:
		commodityGroup.GET("/me/commodities", controller.LoginVerification, controller.QuerySelfCommodities)
		//to publish commodity:
		commodityGroup.POST("/commodities", controller.LoginVerification, controller.PublishCommodity)
		//delete commodity:
		commodityGroup.DELETE("/commodity/:id", controller.LoginVerification, controller.DeleteCommodity)

	}

	collectionGroup := router.Group("")
	{
		//get self collections:
		collectionGroup.GET("/me/collections", controller.LoginVerification, controller.GetSelfCollections)
		//add collection:
		collectionGroup.POST("/me/collections", controller.LoginVerification, controller.AddCollection)
		//delete collection:
		collectionGroup.DELETE("/me/collections", controller.LoginVerification, controller.DeleteCollections)
	}

	fileGroup := router.Group("")
	{
		fileGroup.POST("/pics", controller.LoginVerification, controller.FileUpload)
		fileGroup.GET("/pics/:filename", controller.GetPicture)
	}
	return router
}
