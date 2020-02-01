package mock

import (
	"github.com/sabigara/go-webapp/api"
)

type UserService struct {
	CreateRet      func() (*api.User, error)
	CreateInvoked bool

	GetRet     func() (*api.User, error)
	GetInvoked bool
}

func (us *UserService) Create(name, email string) (*api.User, error) {
	us.CreateInvoked = true
	return us.CreateRet()
}

func (us *UserService) Get(id string) (*api.User, error) {
	us.GetInvoked = true
	return us.GetRet()
}
