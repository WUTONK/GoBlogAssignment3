package handler

import (
	"GinSqlBlog/models"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InitGin(g gin.IRouter) {
	g.POST("/user/login", Login)
	g.GET("/user/info", Info)
}

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

func makeToken() string {
	// 随机生成 32 位哈希
	// 从 md5 修改为使用 crypto/rand 生成token
	// 因为发现了原本的token前后端不一致问题并非因为随机生成源不一致，而是因为在向api发送token时错误的再次使用了makeToken()
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

	c.JSON(http.StatusOK, models.LoginResp{
		Token: token,
	})
	fmt.Printf("已发送token: %s\n", token)

	// 写入token
	if err := os.WriteFile("../tokenList/token.txt", []byte(token), 0666); err != nil {
		log.Fatal(err)
	}

}

func Info(c *gin.Context) {
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
	if authorization == string(token) {
		c.JSON(http.StatusOK, models.InfoResp{
			NickName: "name1<slice>name2<slice>name3",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "传入token与本地token不一致"})
		fmt.Println("err:传入token与本地token不一致")
	}

}
