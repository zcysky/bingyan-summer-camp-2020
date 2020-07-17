package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JWTClaims struct {
	UserID 	uuid.UUID	`json:"user_id"`
	jwt.StandardClaims
}

func NewJWTClaim() *JWTClaims{
	return new(JWTClaims)
}

