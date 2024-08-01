package main

import (
	"log"

	"github.com/bryantang1107/DuriYum/config"
	"github.com/bryantang1107/DuriYum/middleware"
	"github.com/bryantang1107/DuriYum/routes"
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
	// register route
	routes.UserRoute(r)
	routes.PostRoute(r)
	routes.DurianRoute(r)
	routes.OrderRoute(r)

	// checkout - private
	// array boxed/whole indicator, quantity, durian_id

	// add to cart - local storage

	// contact us - public
	// user id, message, name, email, hp
	// send to superadmin email s

	r.Run()
}
