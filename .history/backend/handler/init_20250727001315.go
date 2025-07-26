package handler

import (
	"GinSqlBlog/models"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册路径
func InitGin(g gin.IRouter) {
	g.POST("/user/login", Login)
	g.POST("/user/postModify", postModify)
}

// 校验用户名和密码
func validUserAndPassword(incomingUsername, incomingPassword string) bool {

	fmt.Printf("收到的用户名%s\n", incomingUsername)
	fmt.Printf("收到的密码%s\n", incomingPassword)
	// 密码123456的密文
	passwordHash := "$2a$10$0y3/DmaCYAIRrcPM52TJ6.S/ax2nkl9Avbfp2XS.rATqvWFNIy58G%!"
	username := "kevin"

	usernameValid := false
	passwordValidErr := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(incomingPassword))

	if incomingUsername == username {
		usernameValid = true
		fmt.Println("用户名验证通过")
	}

	if usernameValid && passwordValidErr == nil {
		return true
	} else {
		return false
	}
}

// 随机生成 32位token
func makeToken() string {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(token)
}
func Login(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	valid := validUserAndPassword(req.Username, req.Password)

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "错误的用户名或密码"})
		return
	}

	token := makeToken()

	c.JSON(http.StatusOK, models.LoginRsp{
		Token: token,
	})
	fmt.Printf("已发送token: %s\n", token)

	// 写入token
	if err := os.WriteFile("../tokenList/token.txt", []byte(token), 0666); err != nil {
		log.Fatal(err)
	}

}

// 报文处理
func postModify(c *gin.Context) {
	// mode : get / append / pop / clear

	var req models.SqlReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// ---初始化---
	dsn := "user=postgres password=123456 host=localhost port=5432 dbname=wutonkdb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("成功打开数据库")
	defer db.Close()

	// 变量初始化
	insertSQL := "UPDATE user_context SET context = $1 WHERE username = $2"
	var context string
	var contextSlice []string
	username := req.UserName
	token := req.Token
	mode := req.Mode
	appendText := req.AppendText

	// 验证token
	validResult, err := validToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		fmt.Println("调用验证token函数时发生错误")
		return
	}
	if !validResult {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		fmt.Println("验证token不通过")
		return
	}

	// context初始化
	err = db.QueryRow("SELECT context FROM user_context WHERE username = $1", username).Scan(&context)
	if err != nil {
		// 如果没有用户在数据库中
		if err != sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			fmt.Printf("contxt初始化发生数据库为空以外的错误: %s\n", err)
			return
		}

		fmt.Printf("数据库为空")
	}

	// ---匹配模式---
	switch mode {

	case "get":
		err := db.QueryRow("SELECT context FROM user_context WHERE username = $1", username).Scan(&context)
		if err != nil {
			fmt.Println("数据get失败!")
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Printf("---select context by user '%s'--- \n", username)
		fmt.Println(context)
		c.JSON(http.StatusOK, models.SqlRsp{
			Context: context,
		})

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
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			fmt.Println("数据append失败!")
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			fmt.Println("数据append失败!")
		}
		fmt.Println("数据append成功!")
		return

	case "clear":
		// 向 context 插入空字符串
		_, err = db.Exec(insertSQL, "", username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			fmt.Printf("用户%s context清空失败! \n", username)
		}
		fmt.Printf("用户%s context清空成功! \n", username)
		return

	}
}

// 校验用户token
func validToken(reqtoken string) (bool, error) {

	// --- 文件读取 ---
	localToken, err := os.ReadFile("../tokenList/token.txt")
	if err != nil {
		fmt.Println("token 文件读取失败")
		return false, err
	}
	if string(localToken) == "" {
		fmt.Println("在token缓存中不存在任何token 无法验证")
		return false, err
	}

	fmt.Printf("req_Token:%s\n", reqtoken)
	fmt.Printf("local_Token:%s\n", string(localToken))

	// --- 校验token---
	if reqtoken == string(localToken) {
		fmt.Println("token验证成功")
		return true, nil
	} else {
		fmt.Println("err:传入token与本地token不一致")
		return false, nil
	}

}
