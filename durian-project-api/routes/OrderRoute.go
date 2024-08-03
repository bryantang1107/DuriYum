package routes

import (
	"github.com/bryantang1107/DuriYum/controller"
	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.GET("/orders", controller.GetAllOrders)
	router.GET("/orders/:id", controller.GetOrder)
	router.POST("/orders", controller.CreateOrder)
}
