package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var userHandler *UserHandler

// SetHandlers sets all handlers with their all dependencies injected.
func SetHandlers(user *UserHandler) {
	userHandler = user
}

// Start starts server after settings routes.
func Start(addr string, debug bool) {
	e := echo.New()
	e.HideBanner = true
	if debug {
		e.Debug = true
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/users", userHandler.post)
	e.GET("/users/:id", userHandler.get)

	e.Start(addr)
}
