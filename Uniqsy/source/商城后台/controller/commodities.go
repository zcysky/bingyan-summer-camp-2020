package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
)

//获取所有商品的信息
func GetCommoditiesList(c *gin.Context) {
	//获取前端信息并解析到struct
	var getCommoditiesForm model.GetCommoditiesForm
	err := c.BindJSON(&getCommoditiesForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//根据查询条件从数据库中获取信息
	commodities, err := model.GetCommoditiesInfo(getCommoditiesForm)
	if err != nil {
		fail(c, http.StatusNotFound, err.Error())
	}

	var data []gin.H
	for _, commodity := range commodities {
		data = append(data, gin.H{
			"id":		commodity.ID.String(),
			"title":	commodity.Title,
			"price":	commodity.Price,
			"category":	commodity.Category,
			"picture":	commodity.Picture,
		})
	}
	successHList(c, http.StatusOK, data)
}

//发布新的商品
func PostNewCommodity(c *gin.Context) {
	//获取前端信息并解析到struct
	var postCommodityForm model.PostCommodityForm
	err := c.BindJSON(&postCommodityForm)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//获取经过中间件处理的username
	userName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//在数据中添加商品信息
	err = model.AddNewCommodity(postCommodityForm, userName.(string))
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}
}

//获取当前热词
func GetHotKeyWord(c *gin.Context) {
	keyWords, err := model.GetKeyWords()
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	successStrList(c, http.StatusOK, keyWords)
}