package router

import (
	"github.com/labstack/echo"
	"qq-bot-backend/controller"
)

func initEventGroup(group *echo.Group) {
	group.POST("", controller.EventAdd)
	group.GET("", controller.EventGet)
	group.DELETE("", controller.EventDelete)
	group.PUT("", controller.EventUpdate)
}
