package router

import "github.com/labstack/echo"

func InitRouter(e *echo.Group) {
	eventGroup := e.Group("/event")
	initEventGroup(eventGroup)
}
