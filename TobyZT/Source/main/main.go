package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

func main() {
	ans := rand.Intn(100)
	r := gin.Default()
	r.LoadHTMLFiles("templates/home.html")
	r.GET("/home", func(c *gin.Context) {
		var res string
		fnumber := c.Query("fnumber")
		if fnumber == "" {
			res = ""
			c.HTML(200, "home.html", "HomePage")
		}else {
			num, error := strconv.Atoi(fnumber)
			if error != nil {
				res = "Not a number"
			} else {
				if num == ans {
					res = "Correct!"
				} else if num < ans {
					res = "The answer is bigger."
				} else {
					res = "The answer is smaller"
				}
			}
			c.HTML(200,"home.html", gin.H{
				"res": res,
			})
		}

	})

	r.Run(":3939") // listen and serve on 0.0.0.0:8080
}
