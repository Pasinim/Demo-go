package migrations

import (
	"context"
	"database/sql"
	"demo/utility"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/**
Migration 1 -> testo, se passa m.Up (?) -> Migration 2
Arrange, act assert
*/

func testArticolo(t *testing.T) {
	db := initTestDb()
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
	//time.Sleep(5 * time.Second)
	port, _ := dbContainer.Instance.MappedPort(context.Background(), "5432")
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
	return db
}
