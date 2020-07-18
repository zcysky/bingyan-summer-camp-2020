package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"warmUp/config"
	"warmUp/controller"
)

var e *echo.Echo

func main() {
	e = echo.New()

	g := e.Group("/user")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: config.Config.JWT.JWTSigningMethod,
		SigningKey:    []byte(config.Config.JWT.JWTSecret),
		//Skipper: SkipperFunc,
	}))

	controller.LoginController(e)

	controller.UpdateUserInfoController(g)

	controller.DeleteUserController(g)

	controller.GetUserInfoController(g)

	e.Logger.Fatal(e.Start(":1323"))
}

//func SkipperFunc(ctx echo.Context) bool {
//
//	matched, err := regexp.Match("/user/.*?password=.*$", []byte(ctx.Request().RequestURI))
//	if err != nil{
//		_ =  ctx.String(http.StatusInternalServerError, "regexp error")
//	}
//
//	filter1 := (ctx.Request().Method == http.MethodGet) && matched
//
//	if filter1 {
//		controller.LoginController(e)
//		return true
//	}
//
//
//	return false
//}
