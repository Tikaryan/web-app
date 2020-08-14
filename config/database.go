package config

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://aryan:password@localhost/menu?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
