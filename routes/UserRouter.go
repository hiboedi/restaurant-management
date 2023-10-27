package routes

import (
	controller "restorant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/sign_up", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
