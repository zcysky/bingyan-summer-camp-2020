package controller

import (
	"fmt"
	"net/http"
	"warmup/model"
	"warmup/config"
	"github.com/labstack/echo"
)

func HandleDelete(c echo.Context)error{
	targetId:=c.Param("id")
	claims:=config.JwtToken{}
	err:=ReadJwtToken(c,&claims)
	if err != nil {
		return c.String(http.StatusBadRequest, "无法读取jwt令牌")
	}
	fmt.Println(claims.Uid,targetId)
	if(claims.Type!="admin"){
		return c.String(http.StatusBadRequest, "没有删除权限")
	}
	err=model.DeleteUser(targetId)
	if err != nil {
		return c.String(http.StatusBadRequest, "数据库删除错误")
	}
	return c.String(http.StatusBadRequest, "删除成功")
}