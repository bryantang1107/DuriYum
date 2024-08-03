package routes

import (
	"github.com/bryantang1107/DuriYum/controller"
	"github.com/bryantang1107/DuriYum/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoute(router *gin.Engine) {
	// for Admin
	router.GET("/post", controller.GetAllPosts)
	// for User
	router.GET("/post/:id", middleware.ValidateJWT, controller.GetPost)
	// for User
	router.POST("/post", controller.CreatePost)
}
