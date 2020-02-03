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
	api.UserUsecase
}

// NewUserHandler returns new UserHandler
func NewUserHandler(userUsecase api.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) post(c echo.Context) error {
	m := map[string]string{}
	if err := c.Bind(&m); err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	user, err := h.UserUsecase.Create(m["name"], m["email"])
	if err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	c.JSON(http.StatusCreated, user)
	return nil
}

func (h *UserHandler) get(c echo.Context) error {
	id := c.Param("id")
	user, err := h.UserUsecase.Get(id)
	if err != nil {
		if errors.Is(err, api.ErrResourceNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
	}
	c.JSON(http.StatusOK, user)
	return nil
}
