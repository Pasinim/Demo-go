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
	//time.Sleep(20 * time.Minute)
	return db
}

//func PopulateTestDB(db *sql.DB) {
//	const query = `insert into collezione values (1, 'Collezione 1');
//		insert into collezione values (1, 'Collezione 2');
//		insert into articolo values (1, 'Scarpe', 11, 1),
//		(2, 'Maglia', 22, 1),
//		(3, 'Panta', 33, 1),
//		(4, 'Maglione', 44, 2);
//
//`
//}
