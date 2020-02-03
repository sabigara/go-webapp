package interactor

import (
	"fmt"
	"github.com/sabigara/go-webapp/api"
)

type UserInteractor struct {
	api.UserRepository
}

func NewUserInteractor(userRepository api.UserRepository) *UserInteractor {
	return &UserInteractor{UserRepository: userRepository}
}

func (ui *UserInteractor) Create(name, email string) (*api.User, error) {
	user := api.NewUser(name, email)
	err := ui.UserRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Create: %w", err)
	}
	return user, nil
}

func (ui *UserInteractor) Get(id string) (*api.User, error) {
	user, err := ui.UserRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Get: %w", err)
	}
	return user, err
}
