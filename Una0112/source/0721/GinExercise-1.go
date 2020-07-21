package main

import (
	_ "database/sql"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	_ "net/http"
)

func main() {
	router:=gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"大小姐就是吊坠的")
	})
	router.Run(":0923")
}