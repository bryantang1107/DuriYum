package controller

import (
	"net/http"

	"github.com/bryantang1107/Jom-Fresh/config"
	"github.com/bryantang1107/Jom-Fresh/models"
	"github.com/bryantang1107/Jom-Fresh/utils"
	"github.com/gin-gonic/gin"
)

func GetAllDurians(c *gin.Context) {
	var durian models.Durian
	result := config.DB.Find(&durian)
	if result.Error != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to retrieve durian records", result.Error)
	}
	utils.WriteLog("Record Count :: "+string(result.RowsAffected), "INFO")
}

func AddDurian(c *gin.Context) {
}

func EditDurian(c *gin.Context) {
}

func DeleteDurian(c *gin.Context) {
}
