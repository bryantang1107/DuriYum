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

	var users []models.User
	conditions := map[string]interface{}{"id": id}
	err := dbutils.Select(&users, conditions)
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

	conditions := map[string]interface{}{"id": id}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	err := dbutils.Update(models.User{}, &user, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Edited", err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
