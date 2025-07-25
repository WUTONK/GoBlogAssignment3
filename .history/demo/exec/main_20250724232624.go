package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "user=postgres password=123456 host=localhost port=5432 dbname=wutonkdb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
