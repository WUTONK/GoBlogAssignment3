package main

import (
	"database/sql"
	"fmt"
	"strings"

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
	_, err := db.Exec(insertSQL, "LENA", "")
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

// 查找是否有该用户在数据库中 和该用户context不为空
func isUserContextNotEmpty(db *sql.DB, username string) (bool, bool, error) {
	var context string
	var UserExist, contextExist bool = true, true
	err := db.QueryRow("SELECT context FROM user_context WHERE username = $1", username).Scan(&context)
	if err != nil {
		strErr := fmt.Sprint(err)
		if strErr == "sql: no rows in result set" {
			UserExist = false
			fmt.Println("用户不存在")
		} else {
			fmt.Printf("出现了查询不到用户外的其他错误: %s\n", err)
			return false, false, err
		}
	}

	if context == "" {
		contextExist = false
		fmt.Println("用户没有对应的context")
	}

	return UserExist, contextExist, err
}

// 在context中 追加/弹出一行/清空
func contextModify(db *sql.DB, username string, appendText string, mode string) {
	// mode : append / pop / clear

	// insertSQL := "INSERT INTO user_context (username, context) VALUES ($1, $2)"
	// _, err := db.Exec(insertSQL, "LENA", "")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("数据插入成功！")

	// INSERT INTO user_context (username, context) VALUES ($1, $2)
	insertSQL := "UPDATE user_context SET context = $1 WHERE username = $2"
	var context string
	var contextSlice []string

	err := db.QueryRow("SELECT context FROM user_context WHERE username = $1", username).Scan(&context)
	if err != nil {
		fmt.Printf("contextModify函数SELECT发生错误: %s\n", err)
		return
	}

	switch mode {
	case "append":

		// 标准格式："context1<slice>context2<slice>context3"
		if context == "" {
			context = appendText
			fmt.Printf("(context is empty)context: %s\n", context)
		} else {
			context = context + "<slice>" + appendText
			fmt.Printf("context: %s\n", context)
		}

		// 插入数据
		_, err = db.Exec(insertSQL, context, username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("数据append成功!")
		return

	case "pop":
		contextSlice = strings.Split(context, "<slice>")

		// 弹出最后一个元素
		if len(contextSlice) > 1 {
			contextSlice = contextSlice[:len(contextSlice)-1]
		} else {
			contextSlice = []string{}
		}

		context = strings.Join(contextSlice, "<slice>")
		// 插入数据
		_, err = db.Exec(insertSQL, context, username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("数据append成功!")
		return

	case "clear":
		// 向 context 插入空字符串
		_, err = db.Exec(insertSQL, "", username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("用户%s context清空成功! \n", username)
		return

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

	username := "TESTUSER"

	// 添加唯一约束
	_, err = db.Exec("ALTER TABLE user_context ADD CONSTRAINT user_context_username_key UNIQUE (username);")
	if err != nil {
		fmt.Println("添加唯一约束时出错：", err)
	} else {
		fmt.Println("唯一约束添加成功！")
	}

	// 1. 新建用户（如果不存在）
	insertSQL := "INSERT INTO user_context (username, context) VALUES ($1, $2)"
	_, err = db.Exec(insertSQL, username, "")
	if err != nil {
		fmt.Println("用户可能已存在，忽略插入错误：", err)
	}

	// 2. 追加一条 context
	// contextModify(db, username, "第一次留言", "append")
	// 3. 再追加一条 context
	// contextModify(db, username, "第二次留言", "append")
	// 4. 弹出最后一条 context
	contextModify(db, username, "", "pop")
	// 5. 清空 context
	// contextModify(db, username, "", "clear")

	// 6. 查询 context 并打印
	context, err := selectContext(db, username)
	if err != nil {
		fmt.Println("查询context出错：", err)
	} else {
		fmt.Printf("用户%s的context内容：%s\n", username, context)
	}

	// insert(db)
	// selectAll(db)
	// context, err := selectContext(db, "LENA")
	// fmt.Println(context)
	// fmt.Printf("error %s\n", err)
	// delete(db)
	// isUserContextNotEmpty(db, "LEN")

	// 添加流程: 用户进入 /user/post 页面 -> 发一条post —> 后端查找是否有该用户在数据库中 没有就新建 (如果有)-> 在 context 中追加
	// 删除最后一条流程: 用户进入 /user/post 页面 -> 删除最后一条 —> 后端查找是否有该用户在数据库中 没有就报错 (如果有)-> 在 context 中弹出最后一条
	// 删除全部流程: 用户进入 /user/post 页面 -> 删除全部 —> 后端查找是否有该用户在数据库中 没有就报错 (如果有)-> 往 context写入空
	// 查询流程: 用户进入 /user/posts 页面 -> 后端查找是否有该用户在数据库中且context不为空 空就返回 nil (如果有)-> 将context读取后返回给前端 ->前端解析
	// 需解耦函数：1.查找是否有该用户在数据库中 和该用户context不为空  2.在context中追加和清空  3.将context返回给后端
	// 1√

	// api所需接口 sqlRsq: 'delAll' / 'delLast' / 'add' .  sqlRsp: <string>.
}
