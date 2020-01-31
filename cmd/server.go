package main

import (
	"os"

	"github.com/sabigara/go-webapp/app"
	"github.com/sabigara/go-webapp/app/echo"
	"github.com/sabigara/go-webapp/app/memory"
)

func inject() {
	userRepository := memory.NewUserRepository()
	userInteractor := app.NewUserInteractor(userRepository)
	userHandler := echo.NewUserHandler(userInteractor)
	
	echo.SetHandlers(userHandler)
}

func main() {
	var addr string
	var debug bool
	if val := os.Getenv("DEBUG"); val == "false" {
		debug = false
		addr = "0.0.0.0:1323"
	} else {
		debug = true
		addr = "localhost:1323"
	}

	inject()

	echo.Start(addr, debug)
}
