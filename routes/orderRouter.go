package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mayejacob/restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/orders", controller.GetOrders())
	incomingRoutes.GET("/api/v1/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/api/v1/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/api/v1/orders/:order_id", controller.UpdateOrder())
}
