package controller

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"iElective/config"
	"iElective/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type JwtClaim struct {
	OpenId string `json:"open_id"`
	jwt.StandardClaims
}

//
type LogCredentials struct {
	AppId     string
	AppSecret string
	Code      string
	GrantType string
}

type WeiXinSession struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
}

var JwtConfig middleware.JWTConfig

func Log(ctx echo.Context) error {
	code := ctx.QueryParam("openId")

	credential := LogCredentials{
		AppId:     config.Config.WeiXin.AppId,
		AppSecret: config.Config.WeiXin.AppSecret,
		Code:      code,
		GrantType: "authorization_code",
	}

	weixinSession, err := WeiXinAuth(credential)
	if err != nil {
		return util.HttpErrorHandle(ctx, err)
	}

	jwtToken, err := CreateJwtToken(weixinSession.OpenId, weixinSession.SessionKey)
	if err != nil {
		return util.HttpErrorHandle(ctx, err)
	}

	// set the Jwt config
	JwtConfig.SigningMethod = config.Config.Jwt.JWTSigningMethod
	JwtConfig.SigningKey = []byte(weixinSession.SessionKey)

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": jwtToken,
	})

}

func WeiXinAuth(credentials LogCredentials) (WeiXinSession, error) {

	// set the raw request
	Url, err := url.Parse(config.Config.WeiXin.RequestAddress)
	if err != nil {
		return WeiXinSession{}, err
	}

	// add hte params
	params := url.Values{}

	// add the whole credential into params
	k := reflect.TypeOf(credentials)
	v := reflect.ValueOf(credentials)
	for i := 0; i < k.NumField(); i++ {
		params.Set(k.Field(i).Name, v.Field(i).String())
	}

	// encode the query into url
	Url.RawQuery = params.Encode()

	// request the weixin service to get the weixin session
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return WeiXinSession{}, err
	}
	defer resp.Body.Close()

	// get the request body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeiXinSession{}, err
	}

	// bind the request body into a struct
	var weixinSession WeiXinSession
	err = json.Unmarshal(body, &weixinSession)
	if err != nil {
		return WeiXinSession{}, err
	}

	return weixinSession, nil

}

func CreateJwtToken(openId string, SessionKey string) (string, error) {

	claims := JwtClaim{
		openId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().
				Add(time.Duration(config.Config.Jwt.JWTTokenLife) * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(SessionKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
