package controller

import (
	"JWT/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetToken(c *gin.Context) {
	// 从model生成一个空的claim
	jwtClaim := new(model.JWTClaims)

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
	//获取URL中给出的token，并进行解析
	strToken := c.Query("token")
	claim, err := ParseToken(strToken)

	//分类处理
	if err != nil {
		//解析失败
		ve, ok := err.(*jwt.ValidationError)
		if ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				//token过期
				c.JSON(http.StatusOK, gin.H{
					"token":	"your token is overdue",
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			//无法解析
			"result":	"failed to verify",
		})
	} else {
		//解析成功
		c.JSON(http.StatusOK, gin.H{
			"result":	"verify successfully",
			"id":		claim.UserID,
		})
	}
}