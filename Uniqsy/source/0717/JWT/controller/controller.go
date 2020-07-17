package controller

import (
	"JWT/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetToken(c *gin.Context) {
	// 从model生成一个空的claim
	jwtClaim := model.NewJWTClaim()

	// 对空的claim进行内容填充并生成加密后的token
	singedToken, err := GenToken(jwtClaim)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token":	singedToken,
		})
	}
}

func VerifyToken(c *gin.Context) {
	strToken := c.Query("token")
	claim, err := ParseToken(strToken)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				c.JSON(http.StatusOK, gin.H{
					"token":	"your token is overdue",
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"result":	"failed to verify",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":	"verify successfully",
			"id":		claim.UserID,
		})
	}
}