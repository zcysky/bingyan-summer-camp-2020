/* main program of this project */
package main

import (
	"account/model"
	"account/router"
	"log"
)

// main function
func main() {
	err := model.SetupDatabase()
	if err != nil {
		log.Println(err)
	}

	r := router.InitRouter()
	err = r.Run(":3939")
	if err != nil {
		log.Println(err)
	}
}
