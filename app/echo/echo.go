package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sabigara/go-webapp/app"
)

type UserHandler struct {
	app.UserUsecase
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

// Inject injects dependencies for handlers
func Inject(user *UserHandler) {
	userHandler = user
}

// Start starts server after settings routes
func Start(addr string) {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/users", userHandler.post)
	e.GET("/users/:id", userHandler.get)

	e.Logger.Fatal(e.Start(addr))
}
