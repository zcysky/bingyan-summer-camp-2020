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
	"io"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func User(server *echo.Echo, client *mongo.Client, IsLoggedIn echo.MiddlewareFunc) {
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
			u.Id = CreateRandomNumber(6)
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
	server.GET("/user/:id", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		id := c.Param("id")

		if id == claims.Id {
			res := new(ResUserInfo)
			res.Data = new(UserInfo)

			judge, result := CheckUser(client, "id", claims.Id)
			if judge {
				res.Success = true
				res.Data.Username = result.Username
				res.Data.NickName = result.NickName
				res.Data.Mobile = result.Mobile
				res.Data.Email = result.Email
				res.Data.TotalViewCount = result.TotalViewCount
				res.Data.TotalCollectCount = result.TotalCollectCount
			}else{
				res.Success = false
				res.Error = "数据库出错"
				res.Data = nil
			}
			return c.JSON(http.StatusOK, res)
		}

		res := new(ResUserInfoOther)
		res.Data = new(UserInfoOther)

		judge, result := CheckUser(client, "id", id)
		if judge {
			res.Success = true
			res.Data.NickName = result.NickName
			res.Data.Email = result.Email
			res.Data.TotalViewCount = result.TotalViewCount
			res.Data.TotalCollectCount = result.TotalCollectCount
		}else{
			res.Success = false
			res.Error = "未找到该用户"
			res.Data = nil
		}

		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	return
}

func Commodities(server *echo.Echo, client *mongo.Client, IsLoggedIn echo.MiddlewareFunc) {
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
			fmt.Println(judge)
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
	server.POST("/commodities", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		h := new(GoodInfoALL)
		res := new(ResStr)
		if err := c.Bind(h); err != nil {
			fmt.Println("error")
		}
		res.Success = false
		res.Data = ""
		if h.Category < 1 || h.Category > 8 {
			res.Error = "商品类别不合法"
			return c.JSON(http.StatusOK, res)
		}
		if h.Price < 0 {
			res.Error = "商品价格不合法"
			return c.JSON(http.StatusOK, res)
		}
		judge1, _ := CheckGood(client, "pubuser", claims.Id)
		judge2, _ := CheckGood(client, "title", h.Title)
		if judge1 && judge2 {
			res.Error = "商品重复"
			return c.JSON(http.StatusOK, res)
		}

		h.PubUser = claims.Id
		for {
			h.Id = CreateRandomString(8)
			judge, _ := CheckGood(client, "id", h.Id)
			if judge {
			} else {
				break
			}
		}
		res.Error = ""
		res.Data = "ok"
		err := InsertGood(client, h)
		if err != nil {
			res.Error = "数据库出错"
			res.Data = ""
		}
		res.Success = true
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.GET("/commodity/:id", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		id := c.Param("id")
		res := new(ResGoodInfoDetail)
		res.Data = new(GoodInfoDetail)
		judge, g := CheckGood(client, "id", id)
		if judge {
			res.Data.PubUser = g.PubUser
			res.Data.Title = g.Title
			res.Data.Desc = g.Desc
			res.Data.Category = g.Category
			res.Data.Price = g.Price
			res.Data.Picture = g.Picture
			res.Data.ViewCount = g.ViewCount
			res.Data.CollectCount = g.CollectCount
			res.Success = true
			res.Error = ""
			UpdateGood(client, id, "viewcount")
			UserView(client, claims.Id, "totalviewcount")
		}else{
			res.Success = false
			res.Error = "商品不存在"
			res.Data = nil
		}
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.DELETE("/commodity/:id", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		id := c.Param("id")
		res := new(ResStr)
		res.Success = false
		res.Data = ""

		judge1, result := CheckGood(client, "id", id)
		if judge1 {
			if result.PubUser == claims.Id {
				if DeleteGood(client, id) {
					res.Success = true
					res.Error = ""
					res.Data = "ok"
				}else{
					res.Error = "数据库出错"
				}
			}else{
				res.Error = "非本人发布的商品，无法删除"
			}
		}else{
			res.Error = "未找到该商品"
		}
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	return
}

func Me(server *echo.Echo, client *mongo.Client, IsLoggedIn echo.MiddlewareFunc){
	server.GET("/me", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		res := new(ResUserInfo)
		res.Data = new(UserInfo)

		judge, result := CheckUser(client, "id", claims.Id)
		if judge {
			res.Success = true
			res.Data.Username = result.Username
			res.Data.NickName = result.NickName
			res.Data.Mobile = result.Mobile
			res.Data.Email = result.Email
			res.Data.TotalViewCount = result.TotalViewCount
			res.Data.TotalCollectCount = result.TotalCollectCount
		}else{
			res.Success = false
			res.Error = "数据库出错"
			res.Data = nil
		}
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.POST("/me", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		res := new(ResStr)
		res.Success = false
		res.Data = ""

		u := new(UserInfoAll)
		if err := c.Bind(u); err != nil {
			fmt.Println("error")
		}

		reg := regexp.MustCompile(`[0-9]{11}`)
		str := reg.FindAllString(u.Mobile, -1)
		if !strings.Contains(u.Email, "@") {
			res.Error = "邮箱不合法"
			return c.JSON(http.StatusOK, res)
		} else if str == nil {
			res.Error = "手机号不合法"
			return c.JSON(http.StatusOK, res)
		}

		judge, result := CheckUser(client, "id", claims.Id)
		if judge {
			if u.Password != "" {
				u.Password = Encode(u.Password)
				if !UserUpdate(client, claims.Id, "password", u.Password) {
					res.Error = "数据库出错"
					return c.JSON(http.StatusOK, res)
				}
			}
			if u.NickName != result.NickName {
				if !UserUpdate(client, claims.Id, "nickname", u.NickName) {
					res.Error = "数据库出错"
					return c.JSON(http.StatusOK, res)
				}
			}
			if u.Mobile != result.Mobile {
				if !UserUpdate(client, claims.Id, "mobile", str[0]) {
					res.Error = "数据库出错"
					return c.JSON(http.StatusOK, res)
				}
			}
			if u.Email != result.Email {
				if !UserUpdate(client, claims.Id, "email", u.Email) {
					res.Error = "数据库出错"
					return c.JSON(http.StatusOK, res)
				}
			}
		}
		res.Success = true
		res.Data = "ok"
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.GET("/me/commodities", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		res := new(ResUserGoods)
		result := MyGoodsList(client, claims.Id)

		for _, r := range result {
			var N GoodInfoSim
			N.Id = r.Id
			N.Title = r.Title
			res.Data = append(res.Data, &N)
		}
		res.Success = true
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.GET("/me/collections", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		res := new(ResUserGoods)

		judge, result := CheckUser(client, "id", claims.Id)
		if judge {
			res.Data = result.Collections
		}
		res.Success = true
		res.Error = ""
		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.POST("/me/collections", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		h := new(GoodInfoALL)
		if err := c.Bind(h); err != nil {
			fmt.Println("error")
		}

		res := new(ResStr)
		res.Success = false
		res.Data = ""

		judge1, result1 := CheckGood(client, "id", h.Id)
		if judge1 {
			judge2, result2 := CheckUser(client, "id", claims.Id)
			if judge2 {
				for _, r := range result2.Collections {
					if r.Id == h.Id {
						res.Error = "已收藏"
						return c.JSON(http.StatusOK, res)
					}
				}
				var N GoodInfoSim
				N.Id = h.Id
				N.Title = result1.Title
				result2.Collections = append(result2.Collections, &N)
				result2.TotalCollectCount++
				UserDelete(client, claims.Id)
				err := InsertUser(client, result2)
				if err != nil {
					res.Error = "数据库出错"
					return c.JSON(http.StatusOK, res)
				}
			}
		}else{
			res.Error = "商品不存在"
			return c.JSON(http.StatusOK, res)
		}

		res.Success = true
		res.Data = "ok"

		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	server.DELETE("/me/collections", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		h := new(GoodInfoALL)
		if err := c.Bind(h); err != nil {
			fmt.Println("error")
		}

		res := new(ResStr)
		res.Success = false
		res.Data = ""

		find := false
		judge, result := CheckUser(client, "id", claims.Id)

		if judge {
			for i, temp := range result.Collections {
				if temp.Id == h.Id {
					result.Collections = append(result.Collections[:i], result.Collections[i+1:]...)
					find = true
					break
				}
			}
			if find {
				result.TotalCollectCount--
				UserDelete(client, claims.Id)
				err := InsertUser(client, result)
				if err != nil {
					res.Error = "数据库出错"
					return c.JSON(http.StatusOK, res)
				}
			}else{
				res.Error = "未收藏该商品"
				return c.JSON(http.StatusOK, res)
			}
		}else{
			res.Error = "数据库出错"
			return c.JSON(http.StatusOK, res)
		}

		res.Success = true
		res.Data = "ok"

		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	return
}

func File(server *echo.Echo, client *mongo.Client, IsLoggedIn echo.MiddlewareFunc){
	server.POST("/pics", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		claims.Id = ""

		res := new(ResPics)

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		res.Url = "/pics/" + file.Filename

		return c.JSON(http.StatusOK, res)
	}, IsLoggedIn)
	return
}

func Encode(src string) string {
	m := hmac.New(sha256.New, []byte(Info.Code))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

func CreateRandomNumber(len int) string {
	var container string
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

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}