package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)
import (
	"warmup/config"
	"warmup/model"
)

func HandleLogin(c echo.Context) error {
	userId := c.QueryParam("id")
	userPwd := c.QueryParam("pwd")
	fmt.Println(userId)
	registerInfo, err := model.FindUser(userId)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "查询数据库错误")
	}
	if userPwd != registerInfo.Pwd {
		return c.String(http.StatusBadRequest, "密码错误")
	}
	tokenStr := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(registerInfo.Type)
	claims := config.JwtToken{
		Uid: userId,
		Type:registerInfo.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Config.JWT.TokenDuration)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	tokenStr.Claims = claims
	token, err := tokenStr.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return c.String(http.StatusInternalServerError, "无法创建jwt令牌")
	}
	return c.JSON(http.StatusOK, token)
}
