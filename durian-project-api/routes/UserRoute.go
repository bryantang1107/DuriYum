package routes

import (
	"github.com/bryantang1107/Jom-Fresh/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user/:id", controller.GetUser)
	router.POST("/user", controller.AddUser)
	router.PUT("/user/:id", controller.EditUser)
}
