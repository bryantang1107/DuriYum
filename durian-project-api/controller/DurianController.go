package controller

import (
	"net/http"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
)

func GetAllDurians(c *gin.Context) {
	var durian []models.Durian

	// sort by availability
	err := dbutils.Select(&durian, nil, "")
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	c.IndentedJSON(http.StatusOK, durian)
}

func AddDurian(c *gin.Context) {
	var durian models.Durian

	if err := c.BindJSON(&durian); err != nil {
		utils.WriteLog(err.Error(), "ERROR")
		return
	}

	err := dbutils.Insert(&durian)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to insert durian record", err)
		return
	}
	c.IndentedJSON(http.StatusOK, durian)
}

func EditDurian(c *gin.Context) {
	id := c.Param("id")

	var durian models.Durian

	if err := c.BindJSON(&durian); err != nil {
		utils.WriteLog(err.Error(), "ERROR")
		return
	}

	updates := map[string]interface{}{
		"name":         durian.Name,
		"description":  durian.Description,
		"price_in_kg":  durian.PriceInKG,
		"price_in_box": durian.PriceInBox,
		"is_available": durian.IsAvailable,
	}

	conditions := map[string]interface{}{
		"id": id,
	}

	err := dbutils.Update(&models.Durian{}, updates, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to update durian record", err)
		return
	}
	c.IndentedJSON(http.StatusOK, durian)
}

func DeleteDurian(c *gin.Context) {
	id := c.Param("id")

	var durian models.Durian

	conditions := map[string]interface{}{
		"id": id,
	}
	err := dbutils.Delete(&durian, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to delete durian record", err)
		return
	}

	c.IndentedJSON(http.StatusOK, durian)
}
