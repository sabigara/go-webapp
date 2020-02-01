package main

import (
	"os"

	"github.com/sabigara/go-webapp/api/http"
	"github.com/sabigara/go-webapp/api/memory"
)

func inject() {
	userService := memory.NewUserService()
	userHandler := http.NewUserHandler(userService)

	http.SetHandlers(userHandler)
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

	http.Start(addr, debug)
}
