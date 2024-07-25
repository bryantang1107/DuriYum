package routes

import (
	"github.com/bryantang1107/Jom-Fresh/controller"
	"github.com/gin-gonic/gin"
)

func PostRoute(router *gin.Engine) {
	router.GET("/post", controller.GetAllPosts)
	router.GET("/post/:id", controller.GetPost)
	router.POST("/post", controller.CreatePost)
}
