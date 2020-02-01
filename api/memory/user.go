package memory

import (
	"errors"

	"github.com/sabigara/go-webapp/api"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

var registory = make(map[string]*api.User)

func (us *UserService) Create(name, email string) (*api.User, error) {
	user := api.NewUser(name, email)
	registory[user.ID] = user
	return user, nil
}

func (us *UserService) Get(id string) (*api.User, error) {
	if user, ok := registory[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}
