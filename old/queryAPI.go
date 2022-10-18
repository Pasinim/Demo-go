//package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
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

/*
*
Creo una struttura in cui memorizzare la connessione e associo i metodi GET ad essa
*/
type Repository struct {
	db *sql.DB
}

type EcommerceApi struct {
	repository Repository
}

func getAllCollezioni(r Repository) []Collection {
	result := make([]Collection, 0)
	query := "SELECT * FROM collezione"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var c Collection
		rows.Scan(&c.Id, &c.Name)
		result = append(result, c)
	}
	return result
}

/**
* Restituisce gli articoli della collezione "coll". Se "coll" è zero restituisce tutti gli articoli
 */
func getArticoli(db *sql.DB, coll int) []Item {
	result := make([]Item, 0)
	var query string
	var rows *sql.Rows
	var err error
	if coll == 0 {
		query = "SELECT articolo.id, articolo.name, articolo.sku FROM articolo JOIN collezione ON articolo.collezione_id = collezione.id"
		rows, err = db.Query(query)
	} else {
		query = "SELECT articolo.id, articolo.name, articolo.sku FROM articolo JOIN collezione ON articolo.collezione_id = collezione.id WHERE collezione.id = $1"
		rows, err = db.Query(query, coll)
	}
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var it Item
		rows.Scan(&it.Id, &it.Name, &it.Sku)
		result = append(result, it)
	}
	return result
}

func GETCollezione(e EcommerceApi, id int) Collection {
	query := "SELECT name FROM collezione WHERE id = $1"
	rows, err := e.repository.db.Query(query, id)
	var name string
	rows.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	result := Collection{
		Id:       strconv.Itoa(id),
		Name:     name,
		Articles: nil,
	}
	result.Articles = append(result.Articles, getArticoli(e.repository.db, id)...)
	return result
}

/*
* Creo dei metodi che lavorano direttamebte sulle repository
 */
func (e *EcommerceApi) GETCollezione(c *gin.Context) {
	/** Il contesto trasmette le informazioni come ip ecc */
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, e.repository.GETCollezione())
	} else {
		idInt, _ := strconv.Atoi(id)
		c.JSON(http.StatusOK, getCollezione(db, idInt)) //TODO
	}
}

func ginArticoli(c *gin.Context) {
	connStr := "host=localhost port=5432 user=demo password=demo dbname=demo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, getArticoli(db, 0))
}

func main() {
	connStr := "host=localhost port=5432 user=demo password=demo dbname=demo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Errore init()")
		log.Fatal(err)
	}
	if db.Ping() != nil {
		fmt.Println("Server non raggiungibile")
	}
	r := Repository{db: db}
	api := EcommerceApi{repository: r}

	estate := getCollezione(r, 2)
	fmt.Println("catalogo Estate: ", estate)

	router := gin.Default()
	router.GET("/collezione/", api.GETCollezione)
	router.GET("/collezione/:id", api.GETCollezione)
	router.GET("/articolo", ginArticoli)
	router.Run()

}
