package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"warmUp/module_mapper"
	"warmUp/service"
)

func GetUserInfoController(e *echo.Echo) {

	e.GET("/user/:id", getUserInfo)

	e.GET("/user", getAllUserInfo)

}

func getAllUserInfo(context echo.Context) error {
	user := context.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	hostId := claims["host_id"].(string)

	auth, result, err := service.GetUserInfoService(true, hostId, "")
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return context.String(http.StatusNotFound, "no such resource")
		}

		return context.String(http.StatusInternalServerError, "unknown error")

	}

	if !auth {
		return context.String(http.StatusForbidden, "go away, you are not a admin")
	}

	allUserInfo := result.([]module_mapper.User)

	return context.JSON(http.StatusOK, allUserInfo)
}

func getUserInfo(context echo.Context) error {
	user := context.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	hostId := claims["host_id"].(string)

	auth, result, err := service.GetUserInfoService(false, hostId, "")
	if err != nil {
		if err == mongo.ErrNilDocument {
			return context.String(http.StatusNotFound, "no such resource")
		}

		return context.String(http.StatusInternalServerError, "unknown error")

	}

	if !auth {
		return context.String(http.StatusUnauthorized, "go away, you are not a admin")
	}

	userInfo := result.(module_mapper.User)

	return context.JSON(http.StatusOK, userInfo)
}
