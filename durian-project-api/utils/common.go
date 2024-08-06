package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func WriteLog(message interface{}, msgTyp string) {
	formattedMessage := fmt.Sprintf("%v", message)
	fmt.Println("[ " + msgTyp + " ] :: " + formattedMessage)
}

func HandleErrorResponse(c *gin.Context, status int, message string, err error) {
	c.IndentedJSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}

func CheckModel(model interface{}, expectedKind reflect.Kind) error {
	modelValue := reflect.ValueOf(model)
	WriteLog("Model Value :: "+modelValue.String(), "INFO")
	if modelValue.Kind() != reflect.Ptr {
		return errors.New("model must be a pointer to a struct")
	}
	elem := modelValue.Elem()
	WriteLog("Model Element :: "+elem.String(), "INFO")
	WriteLog("Given Element Kind :: "+elem.Kind().String(), "INFO")
	WriteLog("Expected Element Kind :: "+expectedKind.String(), "INFO")
	if elem.Kind() != expectedKind {
		return errors.New("model must be a pointer to the specified kind")
	}
	return nil
}
