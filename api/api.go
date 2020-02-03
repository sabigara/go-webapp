package api

import (
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
}

type UserUsecase interface {
	Create(name, email string) (*User, error)
	Get(id string) (*User, error)
}

type UserRepository interface {
	Save(*User) error
	Get(id string) (*User, error)
}
