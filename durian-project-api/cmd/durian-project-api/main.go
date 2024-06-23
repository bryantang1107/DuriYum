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

	//login - public

	//signup - public 
	//store fname/lname, email, pw

	//checkout - private
	//array boxed/whole indicator, quantity, durian_id 
	
	//get all durian - public 
	//allow superadmin to amend prices

	//add to cart - local storage

	//contact us - public
	//user id, message, name, email, hp
	//send to superadmin email

	



	r.Run(":10500")
}
