package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"warmUp/module_mapper"
	"warmUp/service"
)

func RegisterController(e *echo.Echo) {
	e.POST("/user", register)

}

func register(context echo.Context) error {

	registerUserInfo := new(module_mapper.RegisterUser)

	if err := context.Bind(registerUserInfo); err != nil{
		return context.String(http.StatusInternalServerError, "bind error")
	}

	done, err := service.RegisterService(*registerUserInfo)
	if err != nil{
		return context.String(http.StatusInternalServerError, "mongodb insert error")
	}

	if !done{
		return context.String(http.StatusUnauthorized, "Fail! wrong register code")
	}

		return context.String(http.StatusOK, "Successful register")
}
