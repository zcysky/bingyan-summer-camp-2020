package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
)

//根据商品的ID获取商品的详细信息
func GetCommodityInfoByID(c *gin.Context) {
	//从URL中获取要查询的商品的ID
	commodityID := c.Param("id")

	//获取商品信息
	commodity, err := model.GetCommodityInfo(commodityID)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//返回查询结果
	data := gin.H{
		"pub_user":      commodity.Publisher,
		"title":         commodity.Title,
		"desc":          commodity.Description,
		"category":      commodity.Category,
		"price":         commodity.Price,
		"picture":       commodity.Picture,
		"view_count":    commodity.View,
		"collect_count": commodity.Collect,
	}
	successH(c, http.StatusOK, data)
}

//根据商品ID删除商品
func DeleteCommodityByID(c *gin.Context) {
	//从URL中获取要查询的商品的ID
	commodityID := c.Param("id")

	//获取经过中间件处理的username
	userName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//删除用户的商品
	err := model.DeleteCommodity(commodityID, userName.(string))
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
	}

	successStr(c, http.StatusNoContent, "ok")
}