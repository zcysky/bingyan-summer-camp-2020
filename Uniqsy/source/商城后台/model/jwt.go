package model

import (
	"github.com/dgrijalva/jwt-go"
	"mall/config"
	"time"
)

//JWT验证的payload
type JWTStruct struct {
	UserName 	string
	jwt.StandardClaims
}

func newJWTClaims() (jwtClaims *JWTStruct) {
	jwtClaims = &JWTStruct{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:	time.Now().Unix(),
			ExpiresAt: 	time.Now().Add(time.Minute *
				time.Duration(config.Config.JWT.EffectiveTime)).Unix(),
		},
	}
	return jwtClaims
}

func GenerateToken(loginForm LoginForm) (tokenStr string, err error) {
	jwtClaims := newJWTClaims()
	jwtClaims.UserName = loginForm.UserName

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenStr, err = token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (userName string, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTStruct{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWT.Secret), nil
		})
	if err != nil {
		return "", err
	}

	return token.Claims.(*JWTStruct).UserName, nil
}