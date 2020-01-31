package app

import (
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUsecase interface {
	Create(name, email string) *User
	Get(id string) *User
}

func NewUser(name, email string) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
}
