package memory

import (
	"errors"
	"github.com/sabigara/go-webapp/app"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var registory = make(map[string]*app.User)

func (userRepo *UserRepository) Save(user *app.User) {
	registory[user.ID] = user
}

func (userRepo *UserRepository) Get(id string) (*app.User, error) {
	if user, ok := registory[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}
