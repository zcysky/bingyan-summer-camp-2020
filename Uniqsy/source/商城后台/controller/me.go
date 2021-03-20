package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
)

//更新自己的个人信息
func UpdateMeInfo(c *gin.Context) {
	//获取信息并转存到struct
	var updateForm model.UpdateForm
	err := c.BindJSON(&updateForm)
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

	//获取原来的用户信息
	user, _, err := model.QueryUser(userName.(string))
	if err != nil {
		fail(c, http.StatusNotFound, err.Error())
		return
	}

	//合并更新内容与原注册信息
	user, err = model.MergeUpdateInfo(updateForm, user)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查新信息的格式是否正确
	err = checkRegisterInfo(user)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//将用户信息更新至数据库
	err = model.UpdateUserInfo(user)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	successStr(c, http.StatusOK, "ok")
}

//获取自己的个人信息
func GetMeInfo(c *gin.Context) {
	//获取经过中间件处理的username
	userName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//获取用户信息
	user, cnt, err := model.QueryUser(userName.(string))
	if err != nil {
		fail(c, http.StatusNotFound, err.Error())
		return
	}

	data := gin.H{
		"username":				user.UserName,
		"nickname":				user.NickName,
		"mobile":				user.Mobile,
		"email":				user.Email,
		"total_view_count":		cnt.ViewCnt,
		"total_collect_count":	cnt.CollectCnt,
	}
	successH(c, http.StatusOK, data)
}

//获取本人上传的商品的信息
func GetMyPostedCommodities	(c *gin.Context) {
	//获取经过中间件处理的username
	userName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//获取商品信息
	commodities, err := model.GetMyCommodities(userName.(string))
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//返回查询信息
	var data []gin.H
	for _, commodity := range commodities {
		data = append(data, gin.H{
			"id":		commodity.ID,
			"title":	commodity.Title,
		})
	}
	successHList(c, http.StatusOK, data)
}

//添加新的收藏
func AddMyNewCollection(c *gin.Context) {
	//获取需要添加的收藏的信息
	var collectionForm model.CollectionForm
	err := c.BindJSON(&collectionForm)
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

	//向数据库中添加收藏信息
	err = model.AddCollection(collectionForm, userName.(string))
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	successStr(c, http.StatusOK, "ok")
}

//获取用户个人的收藏
func GetMyCollectedCommodities(c *gin.Context) {
	//获取经过中间件处理的username
	userName, exists := c.Get("username")
	if !exists {
		fail(c, http.StatusBadRequest, "username is missing")
		return
	}

	//从数据库中获取用户的收藏栏
	collections, err := model.GetUserCollection(userName.(string))
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}

	//返回查询信息
	var data []gin.H
	for _, collection := range collections {
		data = append(data, gin.H{
			"id":		collection.ID,
			"title":	collection.Title,
		})
	}
	successHList(c, http.StatusOK, data)
}

//删除用户的个人收藏
func DeleteMyCollection(c *gin.Context) {
	//获取需要删除的收藏的信息
	var collectionForm model.CollectionForm
	err := c.BindJSON(&collectionForm)
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

	//在数据库中删除收藏
	err = model.DeleteCollection(collectionForm, userName.(string))
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}


}