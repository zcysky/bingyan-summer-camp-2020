package controller

import (
	"2020.7.27/service"
	"2020.7.27/util"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func LoginController(e *echo.Echo) {

	e.GET("/auth/user/:id", login)

}

func login(context echo.Context) error {
	hostId := context.Param("id")
	pwd := context.QueryParam("password")
	pwd = util.Encrypt(pwd)

	isLog, token, err := service.LoginService(hostId, pwd)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return context.String(http.StatusNotFound, "no such a user")
		}
		return context.String(http.StatusInternalServerError, "")
	}

	if !isLog {
		return context.String(http.StatusUnauthorized, "password error")
	}

	return context.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
