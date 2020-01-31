package app

type UserRepository interface {
	Save(*User)
	Get(id string) (*User, error)
}
