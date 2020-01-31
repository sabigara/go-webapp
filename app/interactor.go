package app

type UserInteractor struct {
	UserRepository
}

func NewUserInteractor(userRepository UserRepository) *UserInteractor {
	return &UserInteractor{UserRepository: userRepository}
}

func (interactor *UserInteractor) Create(name, email string) *User {
	user := NewUser(name, email)
	interactor.UserRepository.Save(user)
	return user
}

func (interactor *UserInteractor) Get(id string) *User {
	user, _ := interactor.UserRepository.Get(id)
	return user
}
