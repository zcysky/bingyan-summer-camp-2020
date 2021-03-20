package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"warmup/model"
	"warmup/config"
)

func HandleReadUser(c echo.Context) error {
	claims := config.JwtToken{}
	err := ReadJwtToken(c, &claims)
	if err != nil {
		return c.String(http.StatusBadRequest, "无法读取jwt令牌")
	}
	if claims.Type != "admin" {
		return c.String(http.StatusBadRequest, "没有管理员权限")
	}
	readAll := c.QueryParam("all")
	readId := c.QueryParam("id")
	if readAll == "true" {
		AllUser, err := model.ShowAllUser()
		if err != nil {
			return c.String(http.StatusInternalServerError, "数据库读取错误")
		}
		AllUserJson, err := json.Marshal(AllUser)
		if err != nil {
			return c.String(http.StatusInternalServerError, "json转换错误")
		}
		//fmt.Println(AllUser)
		return c.String(http.StatusOK, string(AllUserJson))

	} else {
		UserInfo, err := model.FindUser(readId)
		UserInfoJson, err := json.Marshal(UserInfo)
		if err != nil {
			return c.String(http.StatusInternalServerError, "json转换错误")
		}
		return c.String(http.StatusOK, string(UserInfoJson))
	}
}
