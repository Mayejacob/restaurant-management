package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mayejacob/restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/tables", controller.GetTables())
	incomingRoutes.GET("/api/v1/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/api/v1/tables", controller.CreateTable())
	incomingRoutes.PATCH("/api/v1/tables/:table_id", controller.UpdateTable())
}
