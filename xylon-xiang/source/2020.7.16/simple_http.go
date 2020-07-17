package main

import (
	"github.com/labstack/echo"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func SimpleHttp(e *echo.Echo) {
	e.GET("/hw", helloWorld)

	e.GET("/rand", randomNum)

	e.POST("/image", saveImage)
}

func saveImage(context echo.Context) error {
	image, err := context.FormFile("image")
	if err != nil {
		return err
	}


	file , err := os.Create(image.Filename)
	if err != nil{
		return err
	}
	defer file.Close()

}

func randomNum(context echo.Context) error {
	seedStr := context.QueryParam("seed")

	seed, err := strconv.ParseInt(seedStr, 10, 64)
	if err != nil {
		return err
	}
	rand.Seed(seed)

	randNum := strconv.FormatInt(rand.Int63(), 10)

	return context.String(http.StatusOK, randNum)
}

func helloWorld(context echo.Context) error {
	return context.String(http.StatusOK, "Hello world")
}