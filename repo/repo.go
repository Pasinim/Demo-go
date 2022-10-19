package repo

import (
	"database/sql"
	"demo/core"
	_ "github.com/lib/pq"
	"log"
)

/*Creo una struttura in cui memorizzare la connessione e associo i metodi GET ad essa*/
type Repository struct {
	db *sql.DB
}

func (r *Repository) GetArticoliCollezioneREPO(idCollezione int) []Item {
	result := make([]core.Item, 0)
	var query string
	var rows *sql.Rows
	var err error
	if idCollezione == 0 {
		query = "SELECT articolo.id, articolo.name, articolo.sku FROM articolo JOIN collezione ON articolo.collezione_id = collezione.id ORDER BY articolo.id"
		rows, err = r.db.Query(query)
	} else {
		query = "SELECT articolo.id, articolo.name, articolo.sku FROM articolo JOIN collezione ON articolo.collezione_id = collezione.id WHERE collezione.id = $1 ORDER BY articolo.id"
		rows, err = r.db.Query(query, idCollezione)
	}
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var it core.Item
		rows.Scan(&it.Id, &it.Name, &it.Sku)
		result = append(result, it)
	}
	return result
}

func (r *Repository) GetArticoliREPO() []core.Item {
	result := make([]core.Item, 0)
	query := "SELECT articolo.id, articolo.name FROM articolo ORDER BY articolo.id"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var it core.Item
		rows.Scan(&it.Id, &it.Name)
		result = append(result, it)
	}
	return result
}

func (r *Repository) GetAllCollezioniREPO() []core.Collection {
	result := make([]core.Collection, 0)
	query := "SELECT * FROM collezione ORDER BY id"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var c core.Collection
		rows.Scan(&c.Id, &c.Name)
		result = append(result, c)
	}
	return result
}

/*
*
Restituisce l'articolo che ha come id 'idArticolo'
*/
func (r *Repository) GetArticoloREPO(idArticolo int) core.Item {
	result := core.Item{
		Id:   0,
		Name: "",
		Sku:  "",
	}

	query := "SELECT id FROM articolo ORDER BY articolo.id"
	ok := false
	var currId int
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&currId)
		if currId == idArticolo {
			ok = true
			break
		}
	}
	if !ok { //???? come restituisco nil
		result.Name = "Articolo non presente"
		return result
	}

	query = "SELECT id, name, sku FROM articolo WHERE id = $1 ORDER BY articolo.id"
	rows, err = r.db.Query(query, currId)
	rows.Next()
	rows.Scan(&result.Id, &result.Name, &result.Sku)
	return result
}

func New() Repository {
	connStr := "host=localhost port=5432 user=demo password=demo dbname=demo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{db: db}
	return r
}
