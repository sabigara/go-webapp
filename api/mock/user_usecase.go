package mock

import (
	"github.com/sabigara/go-webapp/api"
)

type UserInteractor struct {
	CreateRet     func() (*api.User, error)
	CreateInvoked bool

	GetRet     func() (*api.User, error)
	GetInvoked bool
}

func (ui *UserInteractor) Create(name, email, password string) (*api.User, error) {
	ui.CreateInvoked = true
	return ui.CreateRet()
}

func (ui *UserInteractor) GetById(id string) (*api.User, error) {
	ui.GetInvoked = true
	return ui.GetRet()
}
