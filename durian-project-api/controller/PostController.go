package controller

import (
	"errors"
	"net/http"

	"github.com/bryantang1107/DuriYum/dbutils"
	"github.com/bryantang1107/DuriYum/models"
	"github.com/bryantang1107/DuriYum/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPosts(c *gin.Context) {
	var posts models.Post

	err := dbutils.Select(&posts, nil, "created_at desc")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusOK, []interface{}{})
			return
		}
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var posts models.Post
	conditions := map[string]interface{}{"user_id": id}
	err := dbutils.Select(&posts, conditions, "created_at desc")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusOK, []interface{}{})
			return
		}
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.BindJSON(&post); err != nil {
		utils.WriteLog(err.Error(), "ERROR")
		return
	}

	err := dbutils.Insert(&post)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, "Post cannot be created", err)
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}
