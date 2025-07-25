package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func delete(db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM user_context WHERE username=$1")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec("ME")
	if err != nil {
		panic(err)
	}

	fmt.Printf("res = %d", res)
}

func insert(db *sql.DB) {
	insertSQL := "INSERT INTO user_context (username, context) VALUES ($1, $2)"
	_, err := db.Exec(insertSQL, "ME", "HI GIN!")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("数据插入成功！")
}

func selectUsers(db *sql.DB) {
	// 查询用户表
	query := "SELECT username, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	fmt.Println("用户列表：")
	for rows.Next() {
		var username, email string
		err := rows.Scan(&username, &email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Username: %s, Email: %s\n", username, email)
	}
}

func main() {
	dsn := "user=postgres password=123456 host=localhost port=5432 dbname=wutonkdb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("成功打开数据库")
	defer db.Close()

	insert(db)
	selectUsers(db)
	delete(db)

}
