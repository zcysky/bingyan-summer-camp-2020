package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TypeResponse struct {
	Success bool        `json:"success"`
	Hint    string      `json:"hint"`
	Data    interface{} `json:"data"`
}

func MustGetIDFromContext(context echo.Context) string {
	return context.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token).Claims.(jwt.MapClaims)["_id"].(string)
}

func ErrorResponse(context echo.Context, code int, hint string) error {
	return context.JSON(code, TypeResponse{
		Success: false,
		Hint:    hint,
		Data:    nil,
	})
}

func SuccessResponse(context echo.Context, code int, data interface{}) error {
	return context.JSON(code, TypeResponse{
		Success: true,
		Hint:    "",
		Data:    data,
	})
}
