package controller

import (
	"net/http"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	err := dbutils.Select(&posts, nil)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var posts []models.Post
	conditions := map[string]interface{}{"user_id": id}
	err := dbutils.Select(&posts, conditions)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.BindJSON(&post); err != nil {
		utils.WriteLog(err, "ERROR")
		return
	}

	err := dbutils.Insert(&post)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Post cannot be created", err)
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}
