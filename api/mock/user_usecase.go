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

func (us *UserInteractor) Create(name, email string) (*api.User, error) {
	us.CreateInvoked = true
	return us.CreateRet()
}

func (us *UserInteractor) Get(id string) (*api.User, error) {
	us.GetInvoked = true
	return us.GetRet()
}
