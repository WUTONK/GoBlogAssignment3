package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

dsn := "user=postgres password=123456 host=localhost:5432 port=端口号 dbname=数据库名称 sslmode=disable"
db, err := sql.Open("postgres", dsn)
if err != nil {
    panic(err.Error())
}
defer db.Close()
