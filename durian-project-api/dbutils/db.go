package dbutils

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/bryantang1107/Jom-Fresh/config"
	"github.com/bryantang1107/Jom-Fresh/utils"
	"github.com/gin-gonic/gin"
)

func Select(model interface{}, conditions map[string]interface{}) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Slice {
		return errors.New("model must be a pointer to a slice")
	}

	query := config.DB.Where(conditions)

	if err := query.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func Insert(model interface{}, c *gin.Context) (interface{}, error) {
	results := reflect.New(reflect.SliceOf(reflect.TypeOf(model).Elem())).Interface()
	if err := c.BindJSON(results); err != nil {
		return nil, err
	}
	utils.WriteLog("Insert [S]", "INFO")
	result := config.DB.Create(results)
	utils.WriteLog("Insert [E]", "INFO")
	if result.Error != nil {
		return nil, result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")
	return model, nil
}
