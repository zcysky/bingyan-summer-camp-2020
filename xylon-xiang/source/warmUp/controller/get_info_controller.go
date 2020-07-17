package controller

import "github.com/labstack/echo"

func GetUserInfoController(e *echo.Echo) error {

	e.GET("/user/:id", getUserInfo)

	e.GET("/user", getAllUserInfo)

	return nil
}

func getAllUserInfo(context echo.Context) error {

}

func getUserInfo(context echo.Context) error {

}
