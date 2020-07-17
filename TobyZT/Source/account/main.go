/* main program of this project */
package main

import(
	"account/router"

	"log"
)

// main function
func main() {

	r := router.InitRouter()

	err := r.Run(":3939")
	if err != nil {
		log.Println(err)
	}
}
