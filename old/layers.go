package main

/**
Architettura a tre livelli:
1. API
2. ??
3. Query e iterazione con il db
*/

import (
	"database/sql"
)

import (
	_ "github.com/lib/pq"
)

// ----------- LIVELLO 2: chiamate al livello inferiore (?) ----------- //

// ----------- LIVELLO 3: Interazione con il DB ----------- //

/*
Restituisce la lista di tutte le collezioni presenti nel db
@param db database
*/
func getAllCollezioniQUERY(db *sql.DB) []Collection {

}

/*
Restituisce la lista di tutti gli articoli presenti nella collezione idCollezione.
Se idCollezione = 0 restituisce tutti gli articoli in tutte le collezioni
@param db database
*/
func getArticoliCollezioneQUERY(db *sql.DB, idCollezione int) []Item {

	return result
}

/*
Restituisce una lista contenente tutti gli articoli di tutte le collezioni
*/
func getArticoliQUERY(db *sql.DB) []Item {

	return result
}
