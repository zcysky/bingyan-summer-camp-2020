package util

import (
	"github.com/labstack/echo"
)

type TypeResponse struct {
	Success bool        `json:"success"`
	Hint    string      `json:"hint"`
	Data    interface{} `json:"data"`
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
