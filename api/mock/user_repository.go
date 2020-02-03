package mock

import (
	"github.com/sabigara/go-webapp/api"
)

type UserRepository struct {
	SaveRet     func() error
	SaveInvoked bool
	GetRet      func() (*api.User, error)
	GetInvoked  bool
}

func (ur *UserRepository) Save(user *api.User) error {
	ur.SaveInvoked = true
	return ur.SaveRet()
}

func (ur *UserRepository) Get(id string) (*api.User, error) {
	ur.GetInvoked = true
	return ur.GetRet()
}
