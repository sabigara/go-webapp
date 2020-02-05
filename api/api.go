package api

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email string) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
}

type UserUsecase interface {
	Create(name, email, password string) (*User, error)
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(id, name string) (*User, error)
}

type UserRepository interface {
	Save(*User) error
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
}
