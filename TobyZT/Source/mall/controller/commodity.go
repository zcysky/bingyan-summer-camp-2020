package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"net/http"
	"regexp"
)

func GetCommodities(c *gin.Context) {
	var req model.CommodityRequest
	err := c.BindJSON(&req)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	req.Keyword = stripSpecialCharacter(req.Keyword) // to prevent XSS attack
	commodities, total := model.GetCommodities(req)
	var res []gin.H
	for i := 0; i < len(commodities); i++ {
		res = append(res, gin.H{
			"id":       commodities[i].ID.String(),
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
	commodity.Title = stripSpecialCharacter(commodity.Title)
	commodity.Desc = stripSpecialCharacter(commodity.Desc)
	if err != nil {
		failMsg(c, http.StatusBadRequest, err.Error())
		return
	}
	username, exist := c.Get("username")
	if !exist {
		failMsg(c, http.StatusBadRequest, "user not found")
		return
	}
	err = model.AddCommodity(commodity, username.(string))
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

func stripSpecialCharacter(str string) (res string) {
	for _, c := range str {
		reg := regexp.MustCompile("[`@#$%^&*()+=|',.<>/]")
		if !reg.MatchString(string(c)) {
			res += string(c)
		}
	}
	return res
}
