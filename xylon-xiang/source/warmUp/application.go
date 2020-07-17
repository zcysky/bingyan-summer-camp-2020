package main

import (
	"github.com/labstack/echo"
	"warmUp/controller"
)

func main() {
	e := echo.New()

	_ = controller.DeleteUserController(e)

	e.Logger.Fatal(e.Start(":1323"))
}
