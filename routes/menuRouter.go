package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mayejacob/restaurant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/menus", controller.GetMenus())
	incomingRoutes.GET("/api/v1/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/api/v1/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/api/v1/menus/:menu_id", controller.UpdateMenu())
}
