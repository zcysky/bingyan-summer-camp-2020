package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("templates/home.html")
	//r.GET("/home", Guess)  //guessing game
	r.GET("/home", Live)
	r.Run(":3939") // listen and serve
}
