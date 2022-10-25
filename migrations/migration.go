package main

import (
	"context"
	"demo/utility"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

//migrate -path migrations/ -database postgres://demo:demo@localhost:5432/demo?sslmode=disable force 0

func main() {
	db := utility.NewTestDatabase()
	//defer db.Instance.Terminate(context.Background())
	/**
	1. Creare db con test container su cui fare le migrazioni (?)
	Come mi connetto al db??
	*/

	//time.Sleep(5 * time.Second)
	port, _ := db.Instance.MappedPort(context.Background(), "5432")
	//devo mappare la porta 5432 su quella che è stata generata esternamente
	connStr := fmt.Sprintf("postgres://demo:demo@127.0.0.1:%s/demo?sslmode=disable", port)
	m, err := migrate.New(
		"file://migrations/test",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
