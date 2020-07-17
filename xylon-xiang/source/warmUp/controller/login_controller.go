package controller

import "github.com/labstack/echo"

func LoginController(e *echo.Echo) error {

	e.GET("user/:id", login)

	return nil
}

func login(context echo.Context) error {
	id := context.Param("id")

}
