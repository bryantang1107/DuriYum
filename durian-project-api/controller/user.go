package controller

import "github.com/gin-gonic/gin"

func UserController(c *gin.Context){
  c.IndentedJSON(200, gin.H{
    "message" : "pong",
  })
}
