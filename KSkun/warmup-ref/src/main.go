package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"warmup-ref/config"
	"warmup-ref/model"
	"warmup-ref/router"
	"warmup-ref/util"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config.InitConfig()
	model.InitModel()
	util.InitUtil()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	apiGroup := e.Group("/api/v1")
	router.InitRouter(apiGroup)

	log.Fatal(e.Start(config.Config.App.Address))
}
