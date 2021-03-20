package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"task1/config"
	"time"
)

type JWTClaims struct {
	UserID	string
	IsAdmin	bool
	jwt.StandardClaims
}

func NewJWTClaim() *JWTClaims{
	jwtClaim := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Config.JWT.EffectiveTime)).Unix(),
		},
	}

	fmt.Println(jwtClaim.IssuedAt)
	fmt.Println(jwtClaim.ExpiresAt)
	fmt.Println(time.Now().Unix())

	return jwtClaim
}