package memory

import (
	"testing"

	"github.com/sabigara/go-webapp/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService(t *testing.T) {
	assert := assert.New(t)
	expected := &api.User{
		Name:  "sabigara",
		Email: "sabigara@example.com",
	}
	us := NewUserService()
	actual, err := us.Create("sabigara", "sabigara@example.com")
	assert.IsType(expected, actual)
	assert.Equal(expected.Name, actual.Name)
	assert.Equal(expected.Email, actual.Email)
	assert.Nil(err)
}

func TestGetUserService(t *testing.T) {
	assert := assert.New(t)
	expected := &api.User{
		ID:    "id",
		Name:  "sabigara",
		Email: "sabigara@example.com",
	}
	us := NewUserService()
	registory["id"] = expected
	actual, err := us.Get("id")
	assert.IsType(expected, actual)
	assert.Equal(expected.ID, actual.ID)
	assert.Equal(expected.Name, actual.Name)
	assert.Equal(expected.Email, actual.Email)
	assert.Nil(err)
}

func TestGetUserServiceErr(t *testing.T) {
	assert := assert.New(t)
	us := NewUserService()
	actual, err := us.Get("id")
	assert.Nil(actual)
	assert.Equal(err.Error(), "user not found")
}
