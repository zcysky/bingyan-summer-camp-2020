package main

import (
	"2020.7.27/config"
	"2020.7.27/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"regexp"
)

var e *echo.Echo

func main() {

	e = echo.New()

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: config.Config.JWT.JWTSigningMethod,
		SigningKey:    []byte(config.Config.JWT.JWTSecret),
		Skipper:       SkipperFunc,
	}))

	controller.LoginController(e)

	controller.RegisterController(e)

	controller.UpdateUserInfoController(e)

	controller.DeleteUserController(e)

	controller.GetUserInfoController(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func SkipperFunc(ctx echo.Context) bool {
	pwd := ctx.QueryParam("password")
	matched, err := regexp.Match("auth/user/.*", []byte(ctx.Request().RequestURI))
	if err != nil {
		_ = ctx.String(http.StatusInternalServerError, "regexp error")
	}

	filter1 := (ctx.Request().Method == http.MethodGet) && matched && (pwd != "")
	filter2 := ctx.Request().Method == http.MethodPost

	if filter1 || filter2 {
		return true
	}

	return false
}
