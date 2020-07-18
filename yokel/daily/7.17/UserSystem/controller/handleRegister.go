package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"../config"
	"../model"
)


func HandleRegister(c echo.Context) error {
	fmt.Println("fine")
	registerInfo := new(config.RegisterInfo)
	err := c.Bind(registerInfo)
	if err != nil {
		return c.String(http.StatusBadRequest, "无法读取数据")
	}
	tmpUid,err:= uuid.NewV4()
	if(err!=nil){
		return c.String(http.StatusInternalServerError, "无法生成uuid")
	}
	registerInfo.Uid =tmpUid.String()
	registerInfo.UserType="general"
	err=model.InsertNewUser(*registerInfo)
	if(err!=nil){
		return c.String(http.StatusInternalServerError,"无法向数据库添加新用户")
	}
	return c.JSON(http.StatusOK, registerInfo)
}
