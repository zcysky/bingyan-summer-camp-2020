package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JWTClaims struct {
	//token的结构，自定义用户id + 标准结构
	UserID 	uuid.UUID	`json:"user_id"`
	jwt.StandardClaims
}

