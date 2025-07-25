package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "user=postgres password=123456 host=localhost port=5432 dbname=wutonkdb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("成功打开数据库")
	defer db.Close()

	query := "SELECT * FROM user_context"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	fmt.Println("用户列表：")
	for rows.Next() {
		var username, context string
		err := rows.Scan(&username, &context)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Username: %s, Email: %s\n", username, email)
	}

}
