package middlewr

import (
	"2020.7.27/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClaims struct {
	HostId string `json:"host_id"`
	jwt.StandardClaims
}

func CreateJwtToken(hostId string) (string, error) {

	claims := JwtClaims{
		hostId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Config.JWT.JWTTokenLife) * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(config.Config.JWT.JWTSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}
