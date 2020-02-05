package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// authMiddlewareConfig defines the config for Session middleware.
	authMiddlewareConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper      middleware.Skipper
		Authenticate func(c echo.Context) error
	}
)

var (
	// DefaultConfig is the default Session middleware config.
	authMiddlewareDefaultConfig = authMiddlewareConfig{
		Skipper: middleware.DefaultSkipper,
	}
)

func authMiddlewareWithConfig(config authMiddlewareConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = authMiddlewareDefaultConfig.Skipper
	}
	if config.Authenticate == nil {
		panic("Authenticate function is required for authMiddleware")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			if err := config.Authenticate(c); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			return next(c)
		}
	}
}
