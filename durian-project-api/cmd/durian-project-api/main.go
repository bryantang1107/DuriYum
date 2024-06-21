package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/bryantang1107/durian-project-api/middleware"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user": "password",
	}))

	authorized.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":10500")
	fmt.Print("hello")
}
