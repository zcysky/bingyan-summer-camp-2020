package controller

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"warmUp/service"
)

func DeleteUserController(e *echo.Echo) error {

	e.DELETE("/user/:id", deleteUser)

	return nil
}

func deleteUser(context echo.Context) error {
	id := context.Param("id")

	flag, err := service.DeleteUserService(id)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return context.String(http.StatusNotFound, "no such id")
		}
		return context.String(http.StatusNotFound, "")
	}

	if !flag {
		return context.String(http.StatusUnauthorized, "You aren't a admin, and you can't do this")
	}

	return context.String(http.StatusOK, "you have delete the user")

}
