package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
)

func GetSelfCollections(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	commodities := model.GetSelfCollections(username.(string))
	var res []gin.H
	for i := 0; i < len(commodities); i++ {
		res = append(res, gin.H{
			"id":    commodities[i].ID,
			"title": commodities[i].Title,
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error": "", "data": res})
}

func AddCollection(c *gin.Context) {
	var form struct{ id string }
	err := c.BindJSON(&form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	err = model.AddCollection(username.(string), form.id)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	model.AddCollectCounter(form.id, 1)
	c.JSON(http.StatusCreated, gin.H{"success": true, "error": "", "data": "ok"})
}

func DeleteCollections(c *gin.Context) {
	var form struct{ id string }
	err := c.BindJSON(&form)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusUnauthorized, "user not found")
		return
	}
	err = model.DeleteCollection(username.(string), form.id)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	model.AddCollectCounter(form.id, -1)
	c.JSON(http.StatusCreated, gin.H{"success": true, "error": "", "data": "ok"})
}
