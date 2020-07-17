package controller

import (
	"JWT/config"
	"JWT/model"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(strToken string) (*model.JWTClaims, error) {

	token, err := jwt.ParseWithClaims(strToken, &model.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Secret), nil
		})

	if err != nil {
		return nil, err
	} else {
		claim := token.Claims.(*model.JWTClaims)
		return claim, nil
	}
	
}