package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func User(server *echo.Echo, client *mongo.Client) {
	server.POST("/user", func(c echo.Context) error {
		u := new(UserInfoAll)
		res := new(ResStr)
		if err := c.Bind(u); err != nil {
			fmt.Println("error")
		}
		res.Success = false
		res.Data = ""
		reg := regexp.MustCompile(`[0-9]{11}`)
		str := reg.FindAllString(u.Mobile, -1)
		judge, _ := CheckUser(client, "username", u.Username)
		if !strings.Contains(u.Email, "@") {
			res.Error = "邮箱不合法"
			return c.JSON(http.StatusOK, res)
		} else if str == nil {
			res.Error = "手机号不合法"
			return c.JSON(http.StatusOK, res)
		} else if judge {
			res.Error = "用户名已存在"
			return c.JSON(http.StatusOK, res)
		}
		u.Password = Encode(u.Password)
		for {
			u.Id = CreateRandomString(6)
			judge, _ := CheckUser(client, "id", u.Id)
			if judge {
			} else {
				break
			}
		}
		err := InsertUser(client, u)
		if err != nil {
			res.Error = "数据库出错"
			return c.JSON(http.StatusOK, res)
		}
		res.Success = true
		res.Error = ""
		res.Data = "ok"
		return c.JSON(http.StatusOK, res)
	})
	server.POST("/user/login", func(c echo.Context) error {
		u := new(UserInfoAll)
		res := new(ResStr)
		if err := c.Bind(u); err != nil {
			fmt.Println("error")
		}
		u.Password = Encode(u.Password)
		judge, result := CheckUserPass(client, u)
		fmt.Println(judge)
		if judge {
			claims := &jwtCustomClaims{
				result.Username,
				false,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
					Id:        result.Id,
					Subject:   result.Email,
				},
			}
			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Generate encoded token and send it as response.
			tokenstr, err := token.SignedString([]byte(Info.JWTsecret))
			if err != nil {
				return err
			}
			res.Success = true
			res.Error = ""
			res.Data = tokenstr
		} else {
			res.Success = false
			res.Error = "用户验证失败"
			res.Data = ""
		}
		return c.JSON(http.StatusOK, res)
	})
	return
}

func Commodities(server *echo.Echo, client *mongo.Client) {
	server.GET("/commodities", func(c echo.Context) error {
		req := new(ReqGoodsList)
		res := new(ResGoodList)
		res.Success = false
		res.Error = "数据不合法"
		res.Data = nil

		var err error
		req.Page, err = strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, res)
		}
		req.Category, err = strconv.Atoi(c.QueryParam("category"))
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, res)
		}
		req.Limit, err = strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, res)
		}
		req.Keyword = c.QueryParam("keyword")

		if req.Keyword != ""{
			h := new(HotWord)
			h.Keyword = req.Keyword
			judge, _ := CheckHot(client, req.Keyword)
			if !judge {
				err := InsertHot(client, h)
				if err != nil {
					fmt.Println(err)
				}
			}else{
				UpdateHot(client, h)
			}
		}

		if req.Limit < 0 || req.Page < 0 {
			return c.JSON(http.StatusOK, res)
		}
		res.Success = true
		res.Error = ""
		list := GoodList(client, req)
		n := len(list)
		if n <= req.Page*req.Limit {
			res.Error = "商品数不足"
			res.Data = nil
		} else if n-req.Page*req.Limit < req.Limit {
			res.Data = list[req.Page*req.Limit : n]
		} else {
			res.Data = list[req.Page*req.Limit : (req.Page+1)*req.Limit]
		}
		return c.JSON(http.StatusOK, res)
	})
	server.GET("/commodities/hot", func(c echo.Context) error {
		res := new(ResHot)
		h := Hot(client)

		for _, H := range h {
			res.Data = append(res.Data, &H.Keyword)
		}

		return c.JSON(http.StatusOK, res)
	})
	return
}







func Encode(src string) string {
	m := hmac.New(sha256.New, []byte(Info.Code))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

func CreateRandomString(len int) string {
	var container string
	//var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var str = "1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
