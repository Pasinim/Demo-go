package main

import (
	"demo/api"
	"demo/pg_repo"
)

func main() {
	r := pg_repo.New()
	router := api.New(r) //sto passando una repository ""diversa"", che funziona dato che mathca con l'inbterfaccia che ho creato all'interno dell'api
	router.Run("localhost:8080")
}
