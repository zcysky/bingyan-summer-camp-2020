package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"warmUp/module_mapper"
	"warmUp/service"
)

func UpdateUserInfoController(e *echo.Echo) {
	e.PUT("/user/:id", updateUserInfo)

}

func updateUserInfo(context echo.Context) error {
	userInfo := new(module_mapper.User)
	if err := context.Bind(userInfo); err != nil {
		return context.String(http.StatusInternalServerError, "bind error")
	}

	service.UpdateUserInfo(*userInfo)

	return context.String(http.StatusOK, "Have updated")
}
