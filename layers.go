package main

/**
Architettura a tre livelli:
1. API
2. ??
3. Query e iterazione con il db
*/

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

import (
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
Creo una struttura in cui memorizzare la connessione e associo i metodi GET ad essa
*/
type Repository struct {
	db *sql.DB
}

type EcommerceApi struct {
	repository Repository
}

// ----------- LIVELLO 1: API ----------- //
func (e *EcommerceApi) GETCollezione(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, e.repository.getAllCollezioniREPO())
	} else {
		idInt, _ := strconv.Atoi(id)
		c.JSON(http.StatusOK, e.repository.getArticoliCollezioneREPO(idInt))
	}
}

func (e *EcommerceApi) GETArticoli(c *gin.Context) {
	c.JSON(http.StatusOK, e.repository.getArticoliREPO())
}

// ----------- LIVELLO 2: chiamate al livello inferiore (?) ----------- //

func (r *Repository) getAllCollezioniREPO() []Collection {
	return getAllCollezioniQUERY(r.db)
}

//func (r *Repository) getArticoliCollezioneREPO(i int) string {
//	items := getArticoliCollezioneQUERY(r.db, i)
//	var str string
//	for _, v := range items {
//		str += strconv.Itoa(v.Id) + v.Name
//	}
//	return str
//}

func (r *Repository) getArticoliCollezioneREPO(i int) []Item {
	return getArticoliCollezioneQUERY(r.db, i)
}

func (r *Repository) getArticoliREPO() []Item {
	return getArticoliQUERY(r.db)
}

// ----------- LIVELLO 3: Interazione con il DB ----------- //

/*
Restituisce la lista di tutte le collezioni presenti nel db
@param db database
*/
func getAllCollezioniQUERY(db *sql.DB) []Collection {
	result := make([]Collection, 0)
	query := "SELECT * FROM collezione"
	rows, err := db.Query(query)
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

/*
*
Restituisce la lista di tutti gli articoli presenti nella collezione idCollezione.
Se idCollezione = 0 restituisce tutti gli articoli in tutte le collezioni
@param db database
*/
func getArticoliCollezioneQUERY(db *sql.DB, idCollezione int) []Item {
	result := make([]Item, 0)
	var query string
	var rows *sql.Rows
	var err error
	if idCollezione == 0 {
		query = "SELECT articolo.id, articolo.name, articolo.sku FROM articolo JOIN collezione ON articolo.collezione_id = collezione.id"
		rows, err = db.Query(query)
	} else {
		query = "SELECT articolo.id, articolo.name, articolo.sku FROM articolo JOIN collezione ON articolo.collezione_id = collezione.id WHERE collezione.id = $1"
		rows, err = db.Query(query, idCollezione)
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

/*
*
Restituisce una lista contenente tutti gli articoli di tutte le collezioni
*/
func getArticoliQUERY(db *sql.DB) []Item {
	result := make([]Item, 0)
	query := "SELECT articolo.id, articolo.name FROM articolo"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var it Item
		rows.Scan(&it.Id, &it.Name)
		result = append(result, it)
	}
	return result
}

/*
*
Funzione che inizializza il server
*/
func initServer() *gin.Engine {
	connStr := "host=localhost port=5432 user=demo password=demo dbname=demo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{db: db}
	api := EcommerceApi{repository: r}
	router := gin.Default()
	router.GET("/collezione/", api.GETCollezione) //creo gli handler
	router.GET("/collezione/:id", api.GETCollezione)
	router.GET("/articoli/", api.GETArticoli)
	return router
}

func main() {
	router := initServer()
	router.Run("localhost:8080")
}
