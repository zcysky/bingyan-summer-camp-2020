package controller

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"task1/config"
	"task1/model"
)

//生成含有jwt信息的token
func GenerateToken(id string, isAdmin bool) (tokenStr string, err error) {
	jwtClaims := model.NewJWTClaim()
	jwtClaims.UserID = id
	jwtClaims.IsAdmin = isAdmin

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenStr, err = token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

//解析token内容，仅判断是否为管理员
func ParseToken(tokenStr string) (isAdmin bool, err error) {
	if tokenStr == "" || len(tokenStr) < 7 {
		return false, errors.New("token cannot be parsed")
	}
	tokenStr = tokenStr[7:]
	fmt.Println(tokenStr)

	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWT.Secret), nil
		})
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("token is invalid")
	}

	claims := token.Claims.(*model.JWTClaims)
	if claims.IsAdmin {
		return true, nil
	} else {
		return false, nil
	}
}