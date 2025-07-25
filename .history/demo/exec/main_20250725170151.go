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

func selectAll(db *sql.DB) {
	// 查询用户表
	query := "SELECT username, context FROM user_context"
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
		fmt.Printf("Username: %s, Context: %s\n", username, context)
	}
}

func selectContext(db *sql.DB, username string) (string, error) {
	var context string
	err := db.QueryRow("SELECT context FROM user_context WHERE username = $1", username).Scan(&context)
	if err != nil {
		return "", err
	}
	fmt.Printf("---select context by user '%s'--- \n", username)
	return context, nil
}

// func isUserContextNotEmpty(db *sql.DB, username string) (string, error) {
// 	var UserExist, contextExist bool

// }

func main() {
	dsn := "user=postgres password=123456 host=localhost port=5432 dbname=wutonkdb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("成功打开数据库")
	defer db.Close()

	// insert(db)
	selectAll(db)
	context, _ := selectContext(db, "WUTON")
	fmt.Println(context)
	// delete(db)

	// 添加流程: 用户进入 /user/post 页面 -> 发一条post —> 后端查找是否有该用户在数据库中 没有就新建 (如果有)-> 在 context 中追加
	// 删除最后一条流程: 用户进入 /user/post 页面 -> 删除最后一条 —> 后端查找是否有该用户在数据库中 没有就报错 (如果有)-> 在 context 中弹出最后一条
	// 删除全部流程: 用户进入 /user/post 页面 -> 删除全部 —> 后端查找是否有该用户在数据库中 没有就报错 (如果有)-> 往 context写入空
	// 查询流程: 用户进入 /user/posts 页面 -> 后端查找是否有该用户在数据库中且context不为空 空就返回 nil (如果有)-> 将context读取后返回给前端 ->前端解析
	// 需解耦函数：1.查找是否有该用户在数据库中 和该用户context不为空  2.在context中追加和清空  3.将context返回给后端

	// api所需接口 sqlRsq: 'delAll' / 'delLast' / 'add' .  sqlRsp: <string>.
}
