package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mayejacob/restaurant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/v1/invoices", controller.GetInvoices())
	incomingRoutes.GET("/api/v1/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/api/v1/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/api/v1/invoices/:invoice_id", controller.UpdateInvoice())
}
