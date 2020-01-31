package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sabigara/go-webapp/app"
)

// UserHandler implements HandlerFunc methods for User domain.
type UserHandler struct {
	app.UserUsecase
}

// NewUserHandler returns new UserHandler
func NewUserHandler(userUsecase app.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) post(c echo.Context) error {
	m := map[string]string{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	user := h.UserUsecase.Create(m["name"], m["email"])
	c.JSON(http.StatusCreated, user)
	return nil
}

func (h *UserHandler) get(c echo.Context) error {
	id := c.Param("id")
	if user := h.UserUsecase.Get(id); user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return nil
}

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
