package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
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

func (h *UserHandler) create(c echo.Context) error {
	m := map[string]string{}
	if err := c.Bind(&m); err != nil {
		return fmt.Errorf("http create: %w", err)
	}
	user, err := h.UserUsecase.Create(m["name"], m["email"], m["password"])
	if err != nil {
		return fmt.Errorf("http create: %w", err)
	}
	c.JSON(http.StatusCreated, user)
	return nil
}

func (h *UserHandler) login(c echo.Context) error {
	m := map[string]string{}
	if err := c.Bind(&m); err != nil {
		return fmt.Errorf("http login: %w", err)
	}
	email := m["email"]
	passwd := m["password"]
	u, err := h.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("http login: %w", err)
	}
	if u.Password != passwd {
		return c.NoContent(http.StatusUnauthorized)
	}
	sess, _ := session.Get("session", c)
	sess.Values["userid"] = u.ID
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) get(c echo.Context) error {
	id := c.Param("id")
	u, err := h.UserUsecase.GetById(id)
	if err != nil {
		if errors.Is(err, api.ErrResourceNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return fmt.Errorf("%w", err)
	}
	c.JSON(http.StatusOK, u)
	return nil
}

func (h *UserHandler) update(c echo.Context) error {
	id := c.Param("id")
	m := map[string]string{}
	if err := c.Bind(&m); err != nil {
		return fmt.Errorf("http update: %w", err)
	}
	u, err := h.UserUsecase.Update(id, m["name"])
	if err != nil {
		return fmt.Errorf("http update: %w", err)
	}
	c.JSON(http.StatusOK, u)
	return nil
}
