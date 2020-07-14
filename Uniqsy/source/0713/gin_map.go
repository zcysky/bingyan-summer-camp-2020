package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostInfo struct {
	Key 	string `json:"key"`
}

func main() {
	var amap map[string]string
	amap = make(map[string]string)
	amap ["today"] = "2020.07.13"
	amap ["yesterday"] = "2020.07.12"
	amap ["tomorrow"] = "2020.07.14"

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var json PostInfo

		getErr := c.ShouldBind(&json)
		if getErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": getErr.Error()})
		}

		val, ok := amap[json.Key]
		if ok {
			c.JSON(http.StatusOK, gin.H{"result": val})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"result": "nothing"})
		}
	})
	r.Run(":8080")
}