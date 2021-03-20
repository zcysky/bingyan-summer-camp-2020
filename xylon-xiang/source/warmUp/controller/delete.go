package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"warmUp/service"
)

func DeleteUserController(e *echo.Echo) {

	e.DELETE("/user/:id", deleteUser)
}

func deleteUser(context echo.Context) error {
	token := context.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	hostId := claims["host_id"].(string)

	id := context.Param("id")

	flag, err := service.DeleteUserService(hostId, id)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return context.String(http.StatusNotFound, "no such id")
		}
		return context.String(http.StatusNotFound, "")
	}

	if !flag {
		return context.String(http.StatusUnauthorized, "You aren't a admin, and you can't do this")
	}

	return context.String(http.StatusOK, "you have delete the token")

}
