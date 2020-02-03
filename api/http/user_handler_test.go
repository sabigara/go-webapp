package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sabigara/go-webapp/api"
	"github.com/sabigara/go-webapp/api/mock"
	"github.com/stretchr/testify/assert"
)

func generateCtx(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestPostUser(t *testing.T) {
	assert := assert.New(t)

	reqJSON := `{"name":"sabigara","email":"sabigara@example.com"}`
	expected := `{"id":"id","name":"sabigara","email":"sabigara@example.com"}`
	ctx, rec := generateCtx(http.MethodPost, "/", reqJSON)

	user := &api.User{
		ID:    "id",
		Name:  "sabigara",
		Email: "sabigara@example.com",
	}
	ui := &mock.UserInteractor{
		CreateRet: func() (*api.User, error) {
			return user, nil
		},
	}
	uh := NewUserHandler(ui)
	err := uh.post(ctx)
	assert.Nil(err)
	assert.Equal(http.StatusCreated, rec.Code)
	assert.JSONEq(expected, rec.Body.String())
	assert.True(ui.CreateInvoked)
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	expected := `{"id":"id","name":"sabigara","email":"sabigara@example.com"}`
	ctx, rec := generateCtx(http.MethodGet, "/id", "")

	user := &api.User{
		ID:    "id",
		Name:  "sabigara",
		Email: "sabigara@example.com",
	}
	us := &mock.UserInteractor{
		GetRet: func() (*api.User, error) {
			return user, nil
		},
	}
	uh := NewUserHandler(us)
	err := uh.get(ctx)
	assert.Nil(err)
	assert.Equal(http.StatusOK, rec.Code)
	assert.JSONEq(expected, rec.Body.String())
	assert.True(us.GetInvoked)
}
