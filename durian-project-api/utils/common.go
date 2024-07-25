package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func WriteLog(message string, msgTyp string) {
	log.Output(1, "[ "+msgTyp+" ] :: "+message)
}

func HandleErrorResponse(c *gin.Context, status int, message string, err error) {
	c.IndentedJSON(status, gin.H{
		"message": message,
		"error":   err,
	})
}
