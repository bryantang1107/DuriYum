package controller

import (
	"net/http"

	"github.com/bryantang1107/Jom-Fresh/config"
	"github.com/bryantang1107/Jom-Fresh/dbutils"
	"github.com/bryantang1107/Jom-Fresh/models"
	"github.com/bryantang1107/Jom-Fresh/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var users models.User
	conditions := map[string]interface{}{"id": id}
	err := dbutils.Select(&users, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusNotFound, "User Profile Not Found", nil)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	user, err := dbutils.Insert(&models.User{}, c)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Created", err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	id := c.Param("id")

	whereClause := map[string]interface{}{"id": id}
	err := dbutils.Select(&[]models.User{}, whereClause)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusNotFound, "User Profile Not Found", err)
		return
	}

	var user models.User
	if err := c.BindJSON(user); err != nil {
		return
	}

	config.DB.Model(&user).Updates(&user)
	c.IndentedJSON(http.StatusOK, user)
}
