package controller

import (
	"2020.7.27/module_mapper"
	"2020.7.27/service"
	"2020.7.27/util"
	"github.com/labstack/echo"
	"net/http"
)

func UpdateUserInfoController(e *echo.Echo) {
	e.PUT("/user/:id", updateUserInfo)

}

func updateUserInfo(context echo.Context) error {
	userInfo := new(module_mapper.User)
	if err := context.Bind(userInfo); err != nil {
		return context.String(http.StatusInternalServerError, "bind error")
	}

	userInfo.Password = util.Encrypt(userInfo.Password)

	service.UpdateUserInfo(*userInfo)

	return context.String(http.StatusOK, "Have updated")
}
