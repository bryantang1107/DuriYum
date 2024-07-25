package main

import (
	"log"

	"github.com/bryantang1107/Jom-Fresh/config"
	"github.com/bryantang1107/Jom-Fresh/middleware"
	"github.com/bryantang1107/Jom-Fresh/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()

	// to connect DB
	config.Connect()

	r.Use(middleware.CORSMiddleware())
	routes.UserRoute(r)
	r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "pong",
		})
	})

	// checkout - private
	// array boxed/whole indicator, quantity, durian_id

	// add to cart - local storage

	// contact us - public
	// user id, message, name, email, hp
	// send to superadmin email s

	r.Run()
}
