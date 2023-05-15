package main

import (
	"fmt"
	"os"

	"github.com/robertobouses/todo-list-estructurada/internal"
)

func main() {
	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	})
	if err != nil {
		panic(err)
	}
	usrRepo, err := repository.NewUserRepo(db)
	fmt.Println(usrRepo)
	if err != nil {
		panic(err)
	}

	usrApp, err := user.NewUserAppService(usrRepo, jwtSvc)
	fmt.Println(usrApp)
	if err != nil {
		panic(err)
	}

	s := http.NewServer(usrApp, jwtSvc)
	s.Run("8080")
}
