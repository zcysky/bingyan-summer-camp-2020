package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
)

var router = gin.Default()

func InitRouter() {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/", controller.Register)
		userGroup.POST("/login", controller.Login)

		userGroup.GET("/:id", controller.GetOthersInfo)
	}

	commoditiesGroup := router.Group("/commodities", controller.VerifyIdentity)
	{
		commoditiesGroup.GET("/",controller.GetCommoditiesList)
		commoditiesGroup.POST("/", controller.PostNewCommodity)
		commoditiesGroup.GET("/hot", controller.GetHotKeyWord)
	}

	commodityGroup := router.Group("/commodity", controller.VerifyIdentity)
	{
		commodityGroup.GET("/:id", controller.GetCommodityInfoByID)
		commodityGroup.DELETE("/:id", controller.DeleteCommodityByID)
	}

	meGroup := router.Group("/me", controller.VerifyIdentity)
	{
		meGroup.POST("/", controller.UpdateMeInfo)
		meGroup.GET("/", controller.GetMeInfo)

		meGroup.GET("/commodities", controller.GetMyPostedCommodities)

		meGroup.POST("/collections", controller.AddMyNewCollection)
		meGroup.GET("/collections", controller.GetMyCollectedCommodities)
		meGroup.DELETE("/collections", controller.DeleteMyCollection)

	}

	miscGroup := router.Group("", controller.VerifyIdentity)
	{
		miscGroup.POST("/pics", controller.UploadPics)
	}
}

func RunRouter() {
	_ = router.Run(":8080")
}
