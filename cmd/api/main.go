package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	r := gin.Default()
	InitDB()
	SeedUser()
	defer CloseDB()
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

	r.POST("/api/login", func(c *gin.Context) {
		type Req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var req Req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"ok": false, "error": "invalid request"})
			return
		}

		var hash string
		err := DB.QueryRow(c, "SELECT password_hash FROM public.users WHERE email=$1", req.Email).Scan(&hash)
		if err != nil {
			c.JSON(401, gin.H{"ok": false, "error": "wrong email or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
			c.JSON(401, gin.H{"ok": false, "error": "wrong email or password"})
			return
		}

		c.JSON(200, gin.H{"ok": true})
	})
	r.Run(":8080")

}
