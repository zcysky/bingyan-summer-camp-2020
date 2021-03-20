package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"warmup/config"
	"warmup/controller"
)

func main() {
	e := echo.New()
	e.POST("/register-form", controller.HandleRegister)
	e.GET("/token", controller.HandleLogin)

	userJwt := e.Group("/user", middleware.JWT([]byte(config.Config.JWT.Secret)))

	//jwt验证后，还要注意，如果不是管理员，要修改的uid应该与令牌uid一致
	userJwt.PUT("/:id", controller.HandleUpdate)
	userJwt.DELETE("/:id", controller.HandleDelete)
	userJwt.GET("", controller.HandleReadUser)

	e.Start(":1323")
}
