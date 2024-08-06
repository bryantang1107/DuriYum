package routes

import (
	"github.com/bryantang1107/DuriYum/controller"
	"github.com/bryantang1107/DuriYum/middleware"
	"github.com/gin-gonic/gin"
)

func DurianRoute(router *gin.Engine) {
	router.GET("/durians", controller.GetAllDurians)
	router.POST("/durians", middleware.ValidateJWT, controller.AddDurian)
	router.PUT("/durians/:id", middleware.ValidateJWT, controller.EditDurian)
	router.DELETE("/durians/:id", middleware.ValidateJWT, controller.DeleteDurian)
}
