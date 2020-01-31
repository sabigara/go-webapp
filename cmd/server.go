package main

import (
	"github.com/sabigara/go-webapp/app"
	"github.com/sabigara/go-webapp/app/echo"
	"github.com/sabigara/go-webapp/app/memory"
)

func main() {
	userHandler := &echo.UserHandler{
		UserUsecase: &app.UserInteractor{
			UserRepository: &memory.UserRepository{},
		},
	}
	echo.Inject(userHandler)

	echo.Start(":1234")
}
