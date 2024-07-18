package routes

import (
	"github.com/bryantang1107/Jom-Fresh/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
