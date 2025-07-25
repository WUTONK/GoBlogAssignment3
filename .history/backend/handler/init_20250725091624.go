package handler

import "github.com/gin-gonic/gin"

func InitGin(g *gin.IRouter) {
	g.POST("/user/login", Login)
	g.GET("/user/info", Info)
}

func Login(c *gin.Context) {

}

func info(c *gin.Context) {

}
