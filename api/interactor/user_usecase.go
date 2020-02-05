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

func (ui *UserInteractor) Create(name, email, password string) (*api.User, error) {
	user := api.NewUser(name, email)
	user.Password = password
	err := ui.UserRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Create: %w", err)
	}
	return user, nil
}

func (ui *UserInteractor) Update(id, name string) (*api.User, error) {
	user, err := ui.UserRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Update: %w", err)
	}
	user.Name = name
	if err := ui.Save(user); err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Update: %w", err)
	}
	return user, nil
}

func (ui *UserInteractor) GetById(id string) (*api.User, error) {
	user, err := ui.UserRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Get: %w", err)
	}
	return user, err
}

func (ui *UserInteractor) GetByEmail(email string) (*api.User, error) {
	user, err := ui.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("interactor.user_usecase.Get: %w", err)
	}
	return user, err
}
