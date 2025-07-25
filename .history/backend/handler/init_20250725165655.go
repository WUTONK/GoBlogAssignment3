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

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册路径
func InitGin(g gin.IRouter) {
	g.POST("/user/login", Login)
	g.GET("/user/info", Info)
}

// 校验用户
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

// 返回报文
func returnPost(db *sql.DB, username string) (string, error) {

	return "", nil
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

	c.JSON(http.StatusOK, models.LoginResp{
		Token: token,
	})
	fmt.Printf("已发送token: %s\n", token)

	// 写入token
	if err := os.WriteFile("../tokenList/token.txt", []byte(token), 0666); err != nil {
		log.Fatal(err)
	}

}

// 输入报文
func InputPost(c *gin.Context) {
	// 逻辑：如果用户不存在 那么创建这个用户 并写入报文
	// 逻辑：如果用户存在 那么追加报文
	// 逻辑：删除报文
}

func Info(c *gin.Context) {
	// --- 校验逻辑 ---
	authorization := c.GetHeader("Authorization") // 获取头内容

	token, err := os.ReadFile("../tokenList/token.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token 文件读取失败"})
		return
	}
	if string(token) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "在token缓存中不存在任何token 无法验证"})
		fmt.Println("err:在token缓存中不存在任何token 无法验证")
		return
	}

	fmt.Printf("authorization:%s\n", authorization)
	fmt.Printf("token:%s\n", string(token))

	// --- 返回报文 ---
	if authorization == string(token) {
		c.JSON(http.StatusOK, models.InfoResp{
			// NickName: "name1<slice>name2<slice>name3",
			NickName: "",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "传入token与本地token不一致"})
		fmt.Println("err:传入token与本地token不一致")
	}

}
