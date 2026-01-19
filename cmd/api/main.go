package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "../../static")
	r.LoadHTMLGlob("../../templates/*")
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/auth", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "auth.html", nil)
	})
	r.Run(":8080")

}
