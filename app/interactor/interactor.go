package interactor

import (
	"github.com/sabigara/go-webapp/app"
)

type UserInteractor struct {
	app.UserRepository
}

func NewUserInteractor(userRepository app.UserRepository) *UserInteractor {
	return &UserInteractor{UserRepository: userRepository}
}

func (ui *UserInteractor) Create(name, email string) *app.User {
	user := app.NewUser(name, email)
	ui.UserRepository.Save(user)
	return user
}

func (ui *UserInteractor) Get(id string) *app.User {
	user, _ := ui.UserRepository.Get(id)
	return user
}
