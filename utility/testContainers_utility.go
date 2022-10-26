package utility

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
)

func InitTestDb() *sql.DB {
	dbContainer := NewTestContainer()
	db, errOpen := sql.Open("postgres", dbContainer.ConnectionString())
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	port, _ := dbContainer.Instance.MappedPort(context.Background(), "5432")
	connStr := fmt.Sprintf("postgres://demo:demo@127.0.0.1:%d/demo?sslmode=disable", port.Int())
	m, err := migrate.New(
		"file://migrations/test",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	fmt.Println("migrate.UP")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT: ", port)
	return db
}
