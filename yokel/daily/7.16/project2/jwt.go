package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type Claims struct {
	Uid uuid.UUID `json:"uid"`
	jwt.StandardClaims
}

//the data format of jwt payload is json
func GernerateToker(c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)
	//generate uuid
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	//expirationTime:=time.Now().Add(time.Minute * time.Duration(config.Expire))
	claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Expire)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token.Claims = claims
	//fmt.Println(claims,expirationTime,time.Now(),expirationTime.Unix(),time.Now().Unix())
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func TokenCheck(c echo.Context)(uuid.UUID,error){
	cookie,err:=c.Cookie("token")
	if err!=nil {
		if err==http.ErrNoCookie  {
			return uuid.UUID{}, c.String(http.StatusUnauthorized,"can't find cookie of token")
		}
		return uuid.UUID{},c.String(http.StatusBadRequest,"cookie error")
	}
	tokenStr:=cookie.Value
	claims:=Claims{}
	//fmt.Println(tokenStr)

	token,err:=jwt.ParseWithClaims(tokenStr,&claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret),nil
	})
	if err!=nil {
		//fmt.Println(token)
		return uuid.UUID{},c.String(http.StatusBadRequest,"token string error")
	}
	err=token.Claims.Valid()
	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return uuid.UUID{},c.String(http.StatusUnauthorized,"Signature invalid")
		}
		return uuid.UUID{},c.String(http.StatusBadRequest,"token error")
	}
	return claims.Uid,nil
}
