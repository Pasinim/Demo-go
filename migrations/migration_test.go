package migrations

import (
	"context"
	"database/sql"
	"demo/utility"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Item struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Sku  int    `db:"sku"`
}

type Collection struct {
	Id       int
	Name     string
	Articles []Item
}

//func TestDBContainerConnection(t *testing.T) {
//	db, err := initTestDb()
//	assert.NotNil(t, err)
//	assert
//}

func TestInsertCollection(t *testing.T) {
	db, err := initTestDb()
	assert.Nil(t, err)
	const query = `INSERT INTO collezione(id, name) VALUES (1, 'collTest')`
	_, err = db.Exec(query)
	assert.Nil(t, err)
	const selectQuery = `SELECT id, name FROM collezione`
	rows := db.QueryRow(selectQuery)
	var result Collection
	rows.Scan(&result.Id, &result.Name)
	expected := Collection{Id: 1, Name: "collTest"}
	assert.Equal(t, result, expected)
}

func TestInsertArticolo(t *testing.T) {
	db, err := initTestDb()
	assert.Nil(t, err)
	const query = `INSERT INTO articolo(id, nome, sku, collection_id)
				VALUES (1, 'test', 111, 1);`
	_, err = db.Query(query)
	assert.Nil(t, err)
	const selectQuery = `SELECT id, nome, sku FROM articolo WHERE id = 1;`
	rows := db.QueryRow(selectQuery)
	var result Item
	rows.Scan(&result.Id, &result.Name, &result.Sku)
	expected := Item{
		Id:   1,
		Name: "test",
		Sku:  111,
	}
	assert.Equal(t, result, expected)
}

func initTestDb() (*sql.DB, error) {
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
	fmt.Println("PORT: ", port)
	//fmt.Println("time.Sleep...")
	return db, err
}
