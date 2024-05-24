package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mayejacob/restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/users", controllers.GetUsers())
	incomingRoutes.GET("/api/v1/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/api/v1/users/signup", controllers.SignUp())
	incomingRoutes.POST("/api/v1/users/login", controllers.Login())
}
