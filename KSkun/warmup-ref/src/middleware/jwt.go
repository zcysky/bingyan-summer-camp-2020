package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"warmup-ref/config"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte(config.Config.JWT.Secret))
}
