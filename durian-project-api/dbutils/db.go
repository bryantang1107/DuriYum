package dbutils

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/bryantang1107/Jom-Fresh/config"
	"github.com/bryantang1107/Jom-Fresh/utils"
)

func Select(model interface{}, conditions map[string]interface{}) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Slice {
		return errors.New("model must be a pointer to a slice")
	}

	query := config.DB

	for key, value := range conditions {
		query = query.Where(key, value)
	}

	start := time.Now()
	utils.WriteLog("DB Query [S]", "INFO")
	result := query.First(model)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error, "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")

	return nil
}

func Insert(model interface{}) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Struct {
		return errors.New("model must be a pointer to a slice")
	}

	start := time.Now()
	utils.WriteLog("DB Query [S]", "INFO")
	result := config.DB.Create(model)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error, "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")
	return nil
}

func Update(model interface{}, updateModel interface{}, conditions map[string]interface{}) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Slice {
		return errors.New("model must be a pointer to a slice")
	}

	query := config.DB

	for key, value := range conditions {
		query.Where(key+" = ? ", value)
	}

	start := time.Now()
	utils.WriteLog("DB Query [S]", "INFO")
	result := query.Model(model).Updates(updateModel)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error, "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")
	return nil
}
