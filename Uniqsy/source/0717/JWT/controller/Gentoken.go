package controller

import (
	"JWT/config"
	"JWT/model"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

func GenToken(jwtClaim *model.JWTClaims) (string, error) {
	// 随机生成一个UUID
	userID := uuid.NewV4()

	// 根据要求填充claim中的信息
	jwtClaim = &model.JWTClaims{
		UserID:         userID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Config.EffectivTime)).Unix(),
		},
	}


	// 利用配置文件中给出的secret进行加密，完成签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	return token.SignedString([]byte(config.Config.Secret))
}