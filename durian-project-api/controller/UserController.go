package controller

import (
	"net/http"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var users []models.User
	conditions := map[string]interface{}{"id": id}
	err := dbutils.SelectOne(&users, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusNotFound, "User Profile Not Found", err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		utils.WriteLog(err, "ERROR")
		return
	}

	err := dbutils.Insert(&user)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Created", err)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		utils.WriteLog(err, "ERROR")
		return
	}
	updates := map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"phone":    user.Phone,
		"unit":     user.Unit,
		"street":   user.Street,
		"city":     user.City,
		"state":    user.State,
		"postcode": user.PostCode,
		"id":       user.ID,
	}

	conditions := map[string]interface{}{
		"id": id,
	}
	err := dbutils.Update(&models.Durian{}, updates, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Edited", err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
