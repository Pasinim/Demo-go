package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"log"
)

func main() {
	m, err := migrate.New(
		"file://migrations/test",
		"postgres://demo:demo@localhost:5432/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Steps(1)
	if err != nil {
		log.Fatal(err)
	}
}
