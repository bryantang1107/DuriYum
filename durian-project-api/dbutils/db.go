package dbutils

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/bryantang1107/DuriYum/config"
	"github.com/bryantang1107/DuriYum/utils"
	"gorm.io/gorm"
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
	result := query.Find(model)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error, "ERROR")
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")

	return nil
}

func SelectOne(model interface{}, conditions map[string]interface{}) error {
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
		utils.WriteLog(result.Error.Error(), "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")

	return nil
}

func Insert(model interface{}) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Struct {
		return errors.New("updateModel must be a pointer to a struct")
	}

	start := time.Now()
	utils.WriteLog("DB Query [S]", "INFO")
	result := config.DB.Create(model)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error.Error(), "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")
	return nil
}

func Update(model interface{}, updateModel map[string]interface{}, conditions map[string]interface{}) error {
	query := config.DB.Model(model)

	for key, value := range conditions {
		query.Where(key+" = ?", value)
	}

	start := time.Now()
	utils.WriteLog("DB Query [S]", "INFO")
	result := query.Updates(updateModel)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error.Error(), "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")
	return nil
}

func Delete(model interface{}, conditions map[string]interface{}) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Struct {
		return errors.New("model must be a pointer to a slice")
	}

	query := config.DB.Model(model)

	for key, value := range conditions {
		query.Where(key+" = ?", value)
	}

	start := time.Now()
	utils.WriteLog("DB Query [S]", "INFO")
	result := query.Delete(model)
	utils.WriteLog("DB Query [E]", "INFO")
	duration := time.Since(start)
	utils.WriteLog("Operation took :: "+duration.String(), "INFO")

	if result.Error != nil {
		utils.WriteLog(result.Error.Error(), "ERROR")
		return result.Error
	}

	utils.WriteLog("Number of rows affected : "+strconv.Itoa(int(result.RowsAffected)), "INFO")
	return nil
}
