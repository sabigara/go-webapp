package app

type UserUsecase interface {
	Create(name, email string) *User
	Get(id string) *User
}
