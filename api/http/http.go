package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/boj/redistore.v1"
)

var userHandler *UserHandler

// SetHandlers sets all handlers with their all dependencies injected.
func SetHandlers(user *UserHandler) {
	userHandler = user
}

func errorHandler(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		he, ok := err.(*echo.HTTPError)
		if !ok {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
			e.Logger.Error(err.Error())
		}
		code := he.Code
		message := he.Message

		if e.Debug {
			message = err.Error()
		}

		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(he.Code)
			} else {
				err = c.JSON(code, message)
			}
			if err != nil {
				e.Logger.Error(err)
			}
		}
	}
}

func authenticate(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return fmt.Errorf("http authenticate: %w", err)
	}
	if sess.Values["userid"] == nil {
		return fmt.Errorf("http authentication: no user id matched with session id")
	}
	return nil
}

func getCSRFToken(c echo.Context) error {
	c.Response().Header().Set("X-CSRF-Token", csrf.Token(c.Request()))
	c.NoContent(http.StatusOK)
	return nil
}

// Start starts server after settings routes.
func Start(addr string, debug bool) {
	e := echo.New()
	e.HideBanner = true
	if debug {
		e.Debug = true
	}
	e.HTTPErrorHandler = errorHandler(e)

	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err.Error())
	}
	defer store.Close()

	requireLogin := authMiddlewareWithConfig(authMiddlewareConfig{Authenticate: authenticate})
	csrf := echo.WrapMiddleware(
		csrf.Protect(
			[]byte("32-byte-long-auth-key"),
			csrf.Path("/"),
		),
	)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(session.Middleware(store))

	e.GET("/", getCSRFToken, csrf)
	e.POST("/users", userHandler.create, csrf)
	e.POST("/login", userHandler.login)
	e.GET("/users/:id", userHandler.get, requireLogin)
	e.PATCH("/users/:id", userHandler.update, csrf, requireLogin)

	e.Logger.Fatal(e.Start(addr))
}
