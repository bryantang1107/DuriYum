package controller

import (
	"net/http"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	var order models.Order

	err := dbutils.PreLoad(&order, "OrderItems", nil, "created_at desc")
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to retrieve user's order", err)
		return
	}

	// TODO:: pagination

	c.IndentedJSON(http.StatusOK, order)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")

	conditions := map[string]interface{}{
		"user_id": id,
	}

	var order models.Order

	err := dbutils.PreLoad(&order, "OrderItems", conditions, "created_at desc")
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to retrieve user's order", err)
		return
	}

	// TODO:: pagination

	c.IndentedJSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		utils.WriteLog(err.Error(), "ERROR")
		return
	}

	err := dbutils.Insert(&order)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Unable to create order", err)
		return
	}

	c.IndentedJSON(http.StatusOK, order)
}
