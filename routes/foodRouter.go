package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mayejacob/restaurant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/foods", controller.GetFoods())
	incomingRoutes.GET("/api/v1/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/api/v1/foods", controller.CreateFood())
	incomingRoutes.PATCH("/api/v1/foods/:food_id", controller.UpdateFood())
}
