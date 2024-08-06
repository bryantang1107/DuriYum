package routes

import (
	"github.com/bryantang1107/DuriYum/controller"
	"github.com/bryantang1107/DuriYum/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.GET("/orders", middleware.ValidateJWT, controller.GetAllOrders)
	router.GET("/orders/:id", middleware.ValidateJWT, controller.GetOrder)
	router.POST("/orders", middleware.ValidateJWT, controller.CreateOrder)
}
