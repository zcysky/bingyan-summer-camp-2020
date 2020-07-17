package controller

import "github.com/labstack/echo"

func RegisterController(e *echo.Echo) error {
	e.POST("/user", register)

	return nil
}

func register(context echo.Context) error {

	return nil
}
