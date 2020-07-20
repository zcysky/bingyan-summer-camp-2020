package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	//"net/http"
	"warmup/config"
	"warmup/model"
)

func HandleUpdate(c echo.Context) error {
	claims:=config.JwtToken{}
	err:=ReadJwtToken(c,&claims)
	if err != nil {
		return c.String(http.StatusBadRequest, "无法读取jwt令牌")
	}
	targetId := c.Param("id")

	if targetId != claims.Uid && claims.Type!="admin" {
		return c.String(http.StatusBadRequest, "没有修改权限")
	}
	updateInfo := config.RegisterInfo{}
	err = c.Bind(&updateInfo)
	if err != nil {
		return c.String(http.StatusBadRequest, "无法读取数据")
	}
	userInfo, err := model.FindUser(targetId)
	if err != nil {
		return c.String(http.StatusBadRequest, "查询数据库错误")
	}
	userInfo = updateInfo
	userInfo.Uid = targetId
	err = model.UpdateUser(userInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.String(http.StatusBadRequest, "数据库无用户")
		}
		return c.String(http.StatusBadRequest, "数据库更新错误")
	}

	return  c.String(http.StatusOK, "更新成功")
}
