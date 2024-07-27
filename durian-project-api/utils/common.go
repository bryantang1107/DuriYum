package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func WriteLog(message interface{}, msgTyp string) {
	formattedMessage := fmt.Sprintf("%v", message)
	fmt.Println("[ " + msgTyp + " ] :: " + formattedMessage)
}

func HandleErrorResponse(c *gin.Context, status int, message string, err error) {
	c.IndentedJSON(status, gin.H{
		"message": message,
		"error":   err,
	})
}
