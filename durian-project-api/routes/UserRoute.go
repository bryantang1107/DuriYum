package routes

import (
	"github.com/bryantang1107/DuriYum/controller"
	"github.com/bryantang1107/DuriYum/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user", middleware.ValidateJWT, controller.GetUser)
	router.POST("/user", controller.AddUser)
	router.POST("/user/login", controller.UserLogin)
	router.PUT("/user", middleware.ValidateJWT, controller.EditUser)
}
