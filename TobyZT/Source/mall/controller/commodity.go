package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"mall/model"
	"net/http"
	"unicode/utf8"
)

func GetCommodities(c *gin.Context) {
	var req model.CommodityRequest
	err := c.BindJSON(&req)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	req.Keyword = html.EscapeString(req.Keyword) // to prevent XSS attack
	commodities, total := model.GetCommodities(req)
	var res []gin.H
	for i := 0; i < len(commodities); i++ {
		res = append(res, gin.H{
			"id":       commodities[i].ID.Hex(),
			"title":    commodities[i].Title,
			"price":    commodities[i].Price,
			"category": commodities[i].Category,
			"picture":  commodities[i].Picture,
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error": "", "total": total, "data": res})
}

func GetHots(c *gin.Context) {
	keywords := model.GetHots(10)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"data":    keywords,
	})
}

func GetCommodityByID(c *gin.Context) {
	id := c.Param("id")
	model.AddViewCounter(id)
	res := model.GetOneCommodity(id)
	if res.Title == "" {
		failMsg(c, http.StatusNotFound, "commodity not found")
		return
	}
	// create history
	username, _ := c.Get("username")
	if username != "" {
		model.CreateHistory(id, username.(string))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"data": gin.H{
			"pub_user":      res.Publisher,
			"title":         res.Title,
			"desc":          res.Description,
			"category":      res.Category,
			"price":         res.Price,
			"picture":       res.Picture,
			"view_count":    res.View,
			"collect_count": res.Collect,
		},
	})
}

func QuerySelfCommodities(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusBadRequest, "user not found")
		return
	}
	commodities := model.GetSelfCommodities(username.(string))
	var res []gin.H
	for i := 0; i < len(commodities); i++ {
		res = append(res, gin.H{
			"id":    commodities[i].ID,
			"title": commodities[i].Title,
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "error": "", "data": res})
}

func PublishCommodity(c *gin.Context) {
	var commodity model.PublishRequest
	err := c.BindJSON(&commodity)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	err = checkRequest(commodity)
	if err != nil {
		failMsg(c, http.StatusUnauthorized, err.Error())
		return
	}
	commodity.Title = html.EscapeString(commodity.Title)
	commodity.Desc = html.EscapeString(commodity.Desc)

	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusBadRequest, "user not found")
		return
	}

	err = model.AddCommodity(commodity, username.(string))
	c.JSON(http.StatusOK, gin.H{"success": true, "error": "", "data": "ok"})
}

func DeleteCommodity(c *gin.Context) {
	id := c.Param("id")
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusBadRequest, "user not found")
		return
	}
	err := model.DeleteCommodity(id, username.(string))
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok",
	})
}

func checkRequest(s model.PublishRequest) (err error) {
	if s.Title == "" || utf8.RuneCountInString(s.Title) > 40 {
		return fmt.Errorf("invalid title")
	}
	if s.Desc == "" || utf8.RuneCountInString(s.Desc) > 150 {
		return fmt.Errorf("invalid description")
	}
	if s.Category < 1 || s.Category > 9 {
		return fmt.Errorf("invalid category")
	}
	if s.Price <= 0 {
		return fmt.Errorf("invalid price")
	}
	return nil
}
