package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sabigara/go-webapp/api/http"
	"github.com/sabigara/go-webapp/api/mysql"
)

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "user:userpass@tcp(db:3306)/database")
	if err != nil {
		panic(err.Error)
	}
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	return db
}

func inject() {
	db := openDB()
	userService := mysql.NewUserService(db)
	userHandler := http.NewUserHandler(userService)

	http.SetHandlers(userHandler)
}

func main() {
	addr := "0.0.0.0:1323"
	var debug bool
	if val := os.Getenv("DEBUG"); val == "true" {
		debug = true
	} else {
		debug = false
	}
	inject()
	http.Start(addr, debug)
}
