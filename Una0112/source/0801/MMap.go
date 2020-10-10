package main

import (
	_ "bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "os"
)

type Moviename struct {
	Username string `json:"username"`
}

func main() {
	var movieDirector map[string]string
	movieDirector = make(map[string]string)

	movieDirector [ "The Shawshank Redemption" ] = "Frank Darabont"
	movieDirector [ "Farewell My Concubine" ] = "Kaige Chen"
	movieDirector [ "Forrest Gump" ] = "Robert Zemeckis"
	movieDirector [ "Inception" ] = "Christopher Nolan"
	movieDirector [ "Titanic" ] = "James Cameron"

	r:=gin.Default()
	r.POST("/",func(c *gin.Context){
		var st Moviename
		err:=c.BindJSON(&st)
		if (err!=nil){
			log.Fatal(err)
		}
		ans, flag := movieDirector[st.Username]
		fmt.Print(ans)
		if flag {
			c.JSON(http.StatusOK,gin.H{"Its director is ":ans})
		} else {
			c.JSON(http.StatusOK,gin.H{"error":"not found"})
		}
	})
	r.Run(":0923")
}
