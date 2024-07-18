package main

import (
	"github.com/bryantang1107/Jom-Fresh/config"
	"github.com/bryantang1107/Jom-Fresh/middleware"
	"github.com/bryantang1107/Jom-Fresh/routes"
	"github.com/gin-gonic/gin"
)

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

	// login - public

	// signup - public
	// store fname/lname, email, pw

	// checkout - private
	// array boxed/whole indicator, quantity, durian_id

	// get all durian - public
	// allow superadmin to amend prices

	// add to cart - local storage

	// contact us - public
	// user id, message, name, email, hp
	// send to superadmin email s

	r.Run(":10500")
}
