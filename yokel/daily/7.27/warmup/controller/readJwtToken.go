package controller

import (
	"net/http"
	"strings"
	"warmup/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ReadJwtToken(c echo.Context,claims *config.JwtToken)error{
	req := c.Request().Header.Get("authorization")
	splitReq := strings.Split(req, "Bearer ")
	tokenStr := splitReq[1]
	//fmt.Println(tokenStr)
	_, err := jwt.ParseWithClaims(tokenStr,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWT.Secret), nil
	})
	if err != nil {
		//fmt.Println(token)
		return c.String(http.StatusOK, "token string error")
	}
	return nil
}
