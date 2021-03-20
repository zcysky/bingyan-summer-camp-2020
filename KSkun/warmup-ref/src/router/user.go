package router

import (
	"github.com/labstack/echo"
	"warmup-ref/controller"
	middleware2 "warmup-ref/middleware"
)

func initUserGroup(group *echo.Group) {
	group.GET("/token", controller.UserGetToken)
	group.POST("/info", controller.UserRegister)
	group.POST("/verify", controller.UserVerify)

	group.PUT("/info", controller.UserUpdateInfo, middleware2.JWTMiddleware())
	group.DELETE("", controller.UserDelete, middleware2.JWTMiddleware())
	group.GET("/info", controller.UserGetInfo, middleware2.JWTMiddleware())
}
