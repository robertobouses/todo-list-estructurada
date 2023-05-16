package main

import (
	"fmt"
	"os"

	"github.com/robertobouses/todo-list-estructurada/internal"
	"github.com/robertobouses/todo-list-estructurada/user"
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

	fmt.Println(db)

	app := user.NewUserAppService()
	s := NewServer(app)
	s.Run("8080")
}
