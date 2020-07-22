package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"qq-bot-backend/config"
	"qq-bot-backend/model"
	"qq-bot-backend/router"
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

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// access_token check
	e.Use(middleware.KeyAuth(func(key string, context echo.Context) (bool, error) {
		return key == config.Config.App.AccessToken, nil
	}))

	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	apiGroup := e.Group("/api/v1")
	router.InitRouter(apiGroup)

	log.Fatal(e.Start(config.Config.App.Address))
}
