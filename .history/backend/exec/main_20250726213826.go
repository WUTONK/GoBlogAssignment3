package main

import (
	"GinSqlBlog/handler"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	if err := os.WriteFile("../tokenList/token.txt", []byte("tokenFile"), 0666); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "WUTONK")
	})

	handler.InitGin(router)

	log.Println("Backend Server is running on port http://127.0.0.1:8080")
	log.Println("Frontend login page is running on path http://localhost:5173/user/login")

	// 服务器运行端口
	router.Run("127.0.0.1:8080")

}
