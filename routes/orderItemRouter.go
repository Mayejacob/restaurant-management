package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mayejacob/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/api/v1/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/api/v1/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/api/v1/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/api/v1/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
