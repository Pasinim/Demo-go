package main

import (
	"demo/api"
	"demo/repo"
)

func main() {
	r := repo.New()
	router := api.New(r)
	router.Run("localhost:8080")
}
