package main

//https://demo.hedgedoc.org/hawmQYCZQnuoORDs3XOd2Q?edit
import (
	"context"
	"database/sql"
	"demo/utility"
	"fmt"
	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"testing"
)

type Item struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Sku  string `db:"sku"`
}

type Collection struct {
	Id       string
	Name     string
	Articles []Item
}

type demoDBContainer struct {
	instance testcontainers.Container
}

func initDemoDB(ctx context.Context, db sql.DB) error {
	const query = `
   CREATE TABLE IF NOT EXISTS articolo(
       id serial4 primary key not null,
       nome varchar(20), 
       sku varchar(20),
       collezione_id serial4
   );`

	_, err := db.ExecContext(ctx, query)
	return err
}

//
//func truncateDemoDB(ctx context.Context, db sql.DB) error {
//	const query = `TRUNCATE demo.articolo`
//	_, err := db.ExecContext(ctx, query)
//	return err
//}

func TestIntegratonDBInsertSelect(t *testing.T) {
	if testing.Short() {
		t.Skip("skipp")
	}
	ctx := context.Background()
	dbContainer := utility.NewTestContainer()
	defer dbContainer.Close(t)
	db, err := sql.Open("postgres", dbContainer.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	err = initDemoDB(ctx, *db)
	if err != nil {
		log.Fatal(err)
	}

	//aggiungo item
	item := Item{Id: 1, Name: "prova testing", Sku: "111"}
	const insertQuery = `INSERT INTO articolo (id, nome, sku ) values ($1, $2, $3)`
	_, err = db.ExecContext(
		ctx,
		insertQuery,
		item.Id,
		item.Name,
		item.Sku)
	if err != nil {
		t.Fatal(err)
	}

	//select
	savedItem := Item{Id: item.Id}
	const findQuery = `SELECT id, nome, sku FROM articolo where id = $1`
	row := db.QueryRowContext(ctx, findQuery, item.Id)
	err = row.Scan(&savedItem.Id, &savedItem.Name, &savedItem.Sku)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(item, savedItem) {
		t.Fatalf("Gli elementi non corrispondono:\n%s", cmp.Diff(item, savedItem))
	}
}

//func TestJoinDB(t *testing.T) {
//	string
//	query := `SELECT id FROM articolo JOIN collezione ON
//    	collezione.id = articolo.id
//    	ORDER BY articolo.id`
//
//}

func main() {
	fmt.Println("asdadasdad")
}
