package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sabigara/go-webapp/api"
)

// UserHandler implements HandlerFunc methods for User domain.
type UserHandler struct {
	api.UserService
}

// NewUserHandler returns new UserHandler
func NewUserHandler(userService api.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) post(c echo.Context) error {
	m := map[string]string{}
	if err := c.Bind(&m); err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	user, err := h.UserService.Create(m["name"], m["email"])
	if err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	c.JSON(http.StatusCreated, user)
	return nil
}

func (h *UserHandler) get(c echo.Context) error {
	id := c.Param("id")
	user, err := h.UserService.Get(id)
	if err != nil {
		if errors.Is(err, api.ErrResourceNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
	}
	c.JSON(http.StatusOK, user)
	return nil
}
