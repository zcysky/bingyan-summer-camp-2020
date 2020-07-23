package controller

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"net/http"
	"warmUp/module_mapper"
	"warmUp/service"
	"warmUp/util"
)

func RegisterController(e *echo.Echo) {
	e.POST("/user", register)

}

func register(context echo.Context) error {

	registerUserInfo := new(module_mapper.RegisterUser)

	if err := context.Bind(registerUserInfo); err != nil {
		return context.String(http.StatusInternalServerError, "bind error")
	}

	registerUserInfo.Password = util.Encrypt(registerUserInfo.Password)

	done, uuid, token, err := service.RegisterService(registerUserInfo)
	if err != nil {
		if err == redis.Nil {
			return context.String(http.StatusBadRequest, "bo such register code")
		}

		return context.String(http.StatusInternalServerError, "mongodb insert error")
	}

	if !done {
		return context.String(http.StatusUnauthorized, "Fail! wrong register code")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"uuid":  uuid,
		"token": token,
	})
}
