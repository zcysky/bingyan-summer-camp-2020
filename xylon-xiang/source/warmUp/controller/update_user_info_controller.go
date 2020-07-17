package controller

import "github.com/labstack/echo"

func UpdateUserInfoController(e *echo.Echo) error {
	e.PUT("/user/:id", updateUserInfo)

	return nil
}

func updateUserInfo(context echo.Context) error {

	return nil
}
