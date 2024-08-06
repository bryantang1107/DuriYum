package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUser(c *gin.Context) {
	id, _ := c.Get("user")

	conditions := map[string]interface{}{
		"id": id,
	}

	var user models.User
	if err := dbutils.SelectOne(&user, conditions); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Created", err)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		utils.WriteLog(err.Error(), "ERROR")
		return
	}

	if err := dbutils.Insert(&user); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Created", err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	id, _ := c.Get("user")

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		utils.WriteLog(err.Error(), "ERROR")
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
		"postcode": user.Postcode,
	}

	conditions := map[string]interface{}{
		"id": id,
	}

	if err := dbutils.Update(&user, updates, conditions); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "User Profile Cannot be Edited", err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func UserLogin(c *gin.Context) {
	user_id := c.Query("userId")

	conditions := map[string]interface{}{
		"id": user_id,
	}
	var user models.User
	if err := dbutils.SelectOne(&user, conditions); err != nil {
		utils.HandleErrorResponse(c, http.StatusUnauthorized, "User Profile Not Found", err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, "Unable to generate JWT token", err)
		return
	}

	// send it back as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*30, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{})
}
