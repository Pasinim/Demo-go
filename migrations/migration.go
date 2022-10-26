package main

import (
	"context"
	"database/sql"
	"demo/utility"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestArticolo(t *testing.T) {
	db := initTestDb()
	defer db.Close()
	const query = `SELECT id FROM articolo`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var result string
	rows.Scan(&result)
	assert.Equal(t, result, "1")

}

func initTestDb() *sql.DB {
	dbContainer := utility.NewTestContainer()
	db, errOpen := sql.Open("postgres", dbContainer.ConnectionString())
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	//defer db.Instance.Terminate(context.Background())
	/**
	1. Creare db con test container su cui fare le migrazioni
	2. Mappo la porta creata da testContainers sulla porta che ho esposto
	3. Indico la cartella delle migrazioni
	*/
	port, _ := dbContainer.Instance.MappedPort(context.Background(), "5432")
	//devo mappare la porta 5432 su quella che è stata generata esternamente
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
	defer m.Drop()
	fmt.Println("PORT: ", port)
	fmt.Println("time.Sleep...")
	time.Sleep(2 * time.Minute)
	return db
}

func main() {
	initTestDb()

}
