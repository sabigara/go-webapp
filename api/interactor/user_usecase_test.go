package interactor

import (
	"testing"

	"github.com/sabigara/go-webapp/api"
	"github.com/sabigara/go-webapp/api/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserInteractor(t *testing.T) {
	assert := assert.New(t)
	expected := &api.User{
		Name:  "sabigara",
		Email: "sabigara@example.com",
	}
	ur := &mock.UserRepository{
		SaveRet: func() error { return nil },
	}
	us := NewUserInteractor(ur)
	actual, err := us.Create("sabigara", "sabigara@example.com")
	assert.True(ur.SaveInvoked)
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
	ur := &mock.UserRepository{
		GetRet: func() (*api.User, error) { return expected, nil },
	}
	ui := NewUserInteractor(ur)
	actual, err := ui.Get("id")
	assert.True(ur.GetInvoked)
	assert.IsType(expected, actual)
	assert.Equal(expected.ID, actual.ID)
	assert.Equal(expected.Name, actual.Name)
	assert.Equal(expected.Email, actual.Email)
	assert.Nil(err)
}

func TestGetUserServiceErr(t *testing.T) {
	assert := assert.New(t)
	ur := &mock.UserRepository{
		GetRet: func() (*api.User, error) { return nil, api.ErrResourceNotFound },
	}
	ui := NewUserInteractor(ur)
	actual, err := ui.Get("id")
	assert.Nil(actual)
	assert.Equal("interactor.user_usecase.Get: resource not found", err.Error())
}
