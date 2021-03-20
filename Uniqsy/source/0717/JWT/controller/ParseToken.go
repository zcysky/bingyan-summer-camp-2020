package controller

import (
	"JWT/config"
	"JWT/model"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(strToken string) (*model.JWTClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(strToken, &model.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Secret), nil
		})

	//根据是否正常解析进行分类处理
	if err != nil {
		return nil, err
	} else {
		claim := token.Claims.(*model.JWTClaims)
		return claim, nil
	}
	
}