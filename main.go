package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//
//func main() {
//	r := pg_repo.New()
//	a := api.New(r)
//	srv := http.Server{
//		Addr:    ":8080",
//		Handler: a,
//	}
//	srv.ListenAndServe()
//}
