package routes

import (
	"github.com/bryantang1107/DuriYum/controller"
	"github.com/gin-gonic/gin"
)

func DurianRoute(router *gin.Engine) {
	router.GET("/durians", controller.GetAllDurians)
	router.POST("/durians", controller.AddDurian)
	router.PUT("/durians/:id", controller.EditDurian)
	router.DELETE("/durians/:id", controller.DeleteDurian)
}
